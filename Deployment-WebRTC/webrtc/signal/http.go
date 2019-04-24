package signal

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/adaickalavan/Go-Incubate/Deployment-WebRTC/webrtc/handler"
)

// HTTPSDPServer starts a HTTP Server that consumes SDPs
func HTTPSDPServer(port string) chan string {
	s := &sdpServer{sdpChan: make(chan string)}
	s.makeMux()
	go s.runSDPServer(port)
	return s.sdpChan
}

type sdpServer struct {
	recoverCount int
	sdpChan      chan string
	mux          *http.ServeMux
}

func (s *sdpServer) makeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sdp", handlerSDP(s.sdpChan))
	mux.HandleFunc("/join", handlerJoin)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static/"))))
	s.mux = mux
}

func (s *sdpServer) runSDPServer(port string) {
	defer func() {
		s.recoverCount++
		if s.recoverCount > 30 {
			log.Fatal("signal.runSDPServer(): Failed to run")
		}
		if r := recover(); r != nil {
			log.Println("signal.runSDPServer(): PANICKED AND RECOVERED")
			log.Println("Panic:", r)
			go s.runSDPServer(port)
		}
	}()

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

func handlerSDP(sdpChan chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Send client SDP to Golang WebRTC server
		sdpChan <- string(body)

		handler.RespondWithJSON(w, http.StatusAccepted, map[string]string{"Result": "Successfully received client SDP"})
	}
}

func handlerJoin(w http.ResponseWriter, r *http.Request) {
	handler.Push(w, "/static/js/join.js")
	tpl, err := template.ParseFiles("/template/join.html")
	if err != nil {
		log.Printf("\nParse error: %v\n", err)
		handler.RespondWithError(w, http.StatusInternalServerError, "ERROR: Template parse error.")
		return
	}
	handler.Render(w, r, tpl, nil)
}

func handlerPublish(w http.ResponseWriter, r *http.Request) {
	handler.Push(w, "/static/js/publish.js")
	tpl, err := template.ParseFiles("/template/publish.html")
	if err != nil {
		log.Printf("\nParse error: %v\n", err)
		handler.RespondWithError(w, http.StatusInternalServerError, "ERROR: Template parse error.")
		return
	}
	handler.Render(w, r, tpl, nil)
}
