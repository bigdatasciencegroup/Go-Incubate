package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"encoding/json"

	"github.com/joho/godotenv"

	"github.com/adaickalavan/Go-Incubate/Deployment-WebRTC/webrtc/handler"

	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v2"
)

var peerConnectionConfig = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
}

const (
	rtcpPLIInterval = time.Second * 3
)

func init() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Everything below is the pion-WebRTC API, thanks for using it ❤️.
	// Create a MediaEngine object to configure the supported codec
	m := webrtc.MediaEngine{}

	// Setup the codecs you want to use.
	// Only support VP8, this makes our proxying code simpler
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))

	// Create the API object with the MediaEngine
	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	// Run SDP server
	s := newSDPServer(api)
	s.run(os.Getenv("LISTENINGADDR"))
}

type sdpServer struct {
	recoverCount int
	api          *webrtc.API
	pcUpload     map[string]*webrtc.PeerConnection
	pcDownload   map[string]*webrtc.PeerConnection
	localTracks  map[string]*webrtc.Track
	mux          *http.ServeMux
}

func newSDPServer(api *webrtc.API) *sdpServer {
	return &sdpServer{
		api: api,
		pcUpload: make(map[string]*webrtc.PeerConnection),
		pcDownload: make(map[string]*webrtc.PeerConnection),
		localTracks: make(map[string]*webrtc.Track),
	}
}

func (s *sdpServer) makeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sdp", handlerSDP(s))
	mux.HandleFunc("/join", handlerJoin)
	mux.HandleFunc("/publish", handlerPublish)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	s.mux = mux
}

func (s *sdpServer) run(port string) {
	defer func() {
		s.recoverCount++
		if s.recoverCount > 1 {
			log.Fatal("signal.runSDPServer(): Failed to run")
		}
		if r := recover(); r != nil {
			log.Println("signal.runSDPServer(): PANICKED AND RECOVERED")
			log.Println("Panic:", r)
			go s.run(port)
		}
	}()

	s.makeMux()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type message struct{
	Name string   `json:"name"`
	SD webrtc.SessionDescription `json:"sd"`
}

func handlerSDP(s *sdpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var offer message
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&offer); err != nil {
			handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Create a new RTCPeerConnection
		pc, err := s.api.NewPeerConnection(peerConnectionConfig)
		if err != nil {
			panic(err)
		}
	
		switch offer.Name {
		case "Publisher":
			// Store the pc handle
			s.pcUpload[offer.Name] = pc

			// Allow us to receive 1 video track
			if _, err = pc.AddTransceiver(webrtc.RTPCodecTypeVideo); err != nil {
				panic(err)
			}

			// Set a handler for when a new remote track starts
			// Add the incoming track to the list of tracks maintained in the server
			s.localTracks[offer.Name] = addOnTrack(pc)
			log.Println("Offer")
		case "Client":
			if len(s.localTracks) == 0{
				handler.RespondWithError(w, http.StatusInternalServerError, "No local track available for peer connection")
				return
			}
			for _,v := range s.localTracks{
				_, err = pc.AddTrack(v)
				if err != nil {
					handler.RespondWithError(w, http.StatusInternalServerError, "Unable to add local track to peer connection")
					return
				}
				break
			}
			log.Println("Answer")
		default:
			handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Set the remote SessionDescription
		err = pc.SetRemoteDescription(offer.SD)
		if err != nil {
			panic(err)
		}

		// Create answer
		answer, err := pc.CreateAnswer(nil)
		if err != nil {
			panic(err)
		}

		// Sets the LocalDescription, and starts our UDP listeners
		err = pc.SetLocalDescription(answer)
		if err != nil {
			panic(err)
		}

		handler.RespondWithJSON(w, http.StatusAccepted,
			map[string]interface{}{
				"Result": "Successfully received incoming client SDP",
				"SD":   answer,
			})
	}
}

func addOnTrack(pc *webrtc.PeerConnection) *webrtc.Track {
	var localTrack *webrtc.Track
	// Set a handler for when a new remote track starts, this just distributes all our packets
	// to connected peers
	pc.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
		// This can be less wasteful by processing incoming RTCP events, then we would emit a NACK/PLI when a viewer requests it
		go func() {
			ticker := time.NewTicker(rtcpPLIInterval)
			for range ticker.C {
				if rtcpSendErr := pc.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: remoteTrack.SSRC()}}); rtcpSendErr != nil {
					log.Println(rtcpSendErr)
				}
			}
		}()

		// Create a local track, all our SFU clients will be fed via this track
		var newTrackErr error
		localTrack, newTrackErr = pc.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "video", "pion")
		if newTrackErr != nil {
			panic(newTrackErr)
		}

		rtpBuf := make([]byte, 1400)
		for {
			i, readErr := remoteTrack.Read(rtpBuf)
			if readErr != nil {
				panic(readErr)
			}

			// ErrClosedPipe means we don't have any subscribers, this is ok if no peers have connected yet
			if _, err := localTrack.Write(rtpBuf[:i]); err != nil && err != io.ErrClosedPipe {
				panic(err)
			}
		}
	})

	return localTrack
}

func handlerJoin(w http.ResponseWriter, r *http.Request) {
	handler.Push(w, "./static/js/join.js")
	tpl, err := template.ParseFiles("./template/join.html")
	if err != nil {
		log.Printf("\nParse error: %v\n", err)
		handler.RespondWithError(w, http.StatusInternalServerError, "ERROR: Template parse error.")
		return
	}
	handler.Render(w, r, tpl, nil)
}

func handlerPublish(w http.ResponseWriter, r *http.Request) {
	handler.Push(w, "./static/js/publish.js")
	tpl, err := template.ParseFiles("./template/publish.html")
	if err != nil {
		log.Printf("\nParse error: %v\n", err)
		handler.RespondWithError(w, http.StatusInternalServerError, "ERROR: Template parse error.")
		return
	}
	handler.Render(w, r, tpl, nil)
}
