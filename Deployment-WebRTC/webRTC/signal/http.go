package signal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/adaickalavan/Go-Incubate/Deployment-WebRTC/webRTC/handler"
)

var recoverCount int
var sdpChan = make(chan string)

// HTTPSDPServer starts a HTTP Server that consumes SDPs
func HTTPSDPServer(port string) <-chan string {
	go runSDPServer(port)

	return sdpChan
}

func runSDPServer(port string) {
	defer func() {
		recoverCount++
		if recoverCount > 30 {
			log.Fatal("signal.runSDPServer(): Failed to run")
		}
		if r := recover(); r != nil {
			log.Println("signal.runSDPServer():PANICKED AND RECOVERED")
			log.Println("Panic:", r)
			go runSDPServer(port)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/sdp", handlerSDP)
	mux.HandleFunc("/join", handlerJoin)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handlerSDP(w http.ResponseWriter, r *http.Request) {
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

func handlerJoin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
