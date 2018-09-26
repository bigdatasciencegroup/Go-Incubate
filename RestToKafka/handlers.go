package main

import (
	"document"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Shopify/sarama"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/definition/", handlerGetWordByID).Methods("GET")
	// muxRouter.HandleFunc("/definition", handlerGetWord).Methods("GET")
	// muxRouter.HandleFunc("/definition", handlerPostWord).Methods("POST")
	muxRouter.HandleFunc("/definition", handlerPostWord).Methods("POST")
	// muxRouter.HandleFunc("/definition", handlerPutWord).Methods("PUT")
	return muxRouter
}

// func handlerGetWord(w http.ResponseWriter, r *http.Request) {
// 	words, err := dictionary.FindAll()
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, words)
// }

func handlerGetWordByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	word, err := ds[query.Get("word")]
	if err != true {
		respondWithError(w, http.StatusInternalServerError, "Not available")
		return
	}
	respondWithJSON(w, http.StatusOK, word)
}

func handlerPostWord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var word document.Word
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&word); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	word.ID = bson.NewObjectId()

	// Send data to Kafka
	ch := producer.Input()
	wordBytes, _ := json.MarshalIndent(word, "", " ")
	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("TOPICNAME"),
		Key:   sarama.StringEncoder(word.Value),
		Value: sarama.ByteEncoder(wordBytes),
	}

	ch <- msg

	// word.ID = bson.NewObjectId()
	// err := dictionary.Insert(word)

	if _, ok := ds[word.Value]; ok {
		respondWithError(w, http.StatusConflict, "Duplicate entry")
		return
	}

	respondWithJSON(w, http.StatusCreated, word)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

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
