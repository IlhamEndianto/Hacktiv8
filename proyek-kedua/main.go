package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const Address = "localhost:8080"

type ResponseMessage struct {
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/", LandingPage)
	http.HandleFunc("/one", PrintLetterOne)
	log.Fatal(http.ListenAndServe(Address, nil))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello World"))
	}
}

func PrintLetterOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responseMessage := ResponseMessage{"1"}
		responseInByte, _ := json.Marshal(responseMessage)
		w.Write(responseInByte)
	}
}
