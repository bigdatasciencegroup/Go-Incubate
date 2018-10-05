package main

import (
	"document"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/definition/", handlerGetWordByID).Methods("GET")
	muxRouter.HandleFunc("/definition", handlerGetWord).Methods("GET")
	muxRouter.HandleFunc("/definition", handlerPostWord).Methods("POST")
	muxRouter.HandleFunc("/definition", handlerPostWord).Methods("POST")
	// muxRouter.HandleFunc("/definition", handlerPutWord).Methods("PUT")
	return muxRouter
}

func handlerGetWord(w http.ResponseWriter, r *http.Request) {
	words, err := dictionary.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, words)
}

func handlerGetWordByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	word, err := dictionary.FindByValue(query.Get("word"))
	// word, err := ds[query.Get("word")]
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, word)
}

func handlerPostWord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//Decode the received post
	word := &document.Word{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(word)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	word.ID = bson.NewObjectId()

	//Prepare message to be sent to Kafka
	wordBytes, _ := json.Marshal(*word)
	msg := &sarama.ProducerMessage{
		Topic:     os.Getenv("TOPICNAMEPOST"),
		Key:       sarama.StringEncoder(word.Value),
		Value:     sarama.ByteEncoder(wordBytes),
		Timestamp: time.Now(),
	}

	//Send message into Kafka queue
	producer.Input() <- msg

	//Verify if the database entry is a duplicate
	// if _, ok := ds[word.Value]; ok {
	// 	respondWithError(w, http.StatusConflict, "Duplicate entry")
	// 	return
	// }

	respondWithJSON(w, http.StatusCreated, word)
}

//Generic wrapper function to respond with error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//Generic wrapper function to respond with JSON message
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		http.Error(w, "HTTP 500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}
