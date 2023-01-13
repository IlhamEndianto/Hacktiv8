package main

import (
	"log"
	"net/http"
)

const Address = "localhost:8080"

func main() {

	http.HandleFunc("/", LandingPage)
	log.Fatal(http.ListenAndServe(Address, nil))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello World"))
	}
}
