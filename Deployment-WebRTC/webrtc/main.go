package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"html/template"
	"time"

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

	// localTrack := <-localTrackChan
	// for {
	// 	log.Println("")
	// 	log.Println("Curl an base64 SDP to start sendonly peer connection")

	// 	recvOnlyOffer := webrtc.SessionDescription{}
	// 	signal.Decode(<-sdpChan, &recvOnlyOffer)

	// 	// Create a new PeerConnection
	// 	peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	_, err = peerConnection.AddTrack(localTrack)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Set the remote SessionDescription
	// 	err = peerConnection.SetRemoteDescription(recvOnlyOffer)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Create answer
	// 	answer, err := peerConnection.CreateAnswer(nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Sets the LocalDescription, and starts our UDP listeners
	// 	err = peerConnection.SetLocalDescription(answer)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Get the LocalDescription and take it to base64 so we can paste in browser
	// 	log.Println(signal.Encode(answer))
	// }
}

type sdpServer struct {
	recoverCount int
	api          *webrtc.API
	pcUpload     []*webrtc.PeerConnection
	pcDownload   []*webrtc.PeerConnection
	mux          *http.ServeMux
}

func newSDPServer(api *webrtc.API) *sdpServer {
	return &sdpServer{api: api}
}

func (s *sdpServer) makeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sdp/", handlerSDP(s))
	mux.HandleFunc("/join", handlerJoin)
	mux.HandleFunc("/publish", handlerPublish)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	s.mux = mux
}

func (s *sdpServer) run(port string) {
	defer func() {
		s.recoverCount++
		if s.recoverCount > 30 {
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

func handlerSDP(s *sdpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		// 	return
		// }

		// bodyStr := string(body)
		// if !isJSON(bodyStr) {
		// 	handler.RespondWithError(w, http.StatusBadRequest, "ERROR: SDP error.")
		// 	return
		// }
		// Send client SDP to Golang WebRTC server
		// sdpChan <- bodyStr

		// offer := webrtc.SessionDescription{}
		// signal.Decode(<-sdpChan, &offer)
		log.Println("Received first offer")

		// Create a new RTCPeerConnection
		pc, err := s.api.NewPeerConnection(peerConnectionConfig)
		if err != nil {
			panic(err)
		}

		// Store the pc handle
		s.pcUpload = append(s.pcUpload, pc)

		// Allow us to receive 1 video track
		if _, err = pc.AddTransceiver(webrtc.RTPCodecTypeVideo); err != nil {
			panic(err)
		}

		// localTrack := addOnTrack(pc)
		_ = addOnTrack(pc)

		// // Set the remote SessionDescription
		// err = pc.SetRemoteDescription(offer)
		// if err != nil {
		// 	panic(err)
		// }

		// // Create answer
		// answer, err := pc.CreateAnswer(nil)
		// if err != nil {
		// 	panic(err)
		// }

		// // Sets the LocalDescription, and starts our UDP listeners
		// err = pc.SetLocalDescription(answer)
		// if err != nil {
		// 	panic(err)
		// }

		answer := "yes this is your answer"
		handler.RespondWithJSON(w, http.StatusAccepted,
			map[string]string{
				"Result":    "Successfully received client SDP",
				"ServerSDP": answer,
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

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
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
