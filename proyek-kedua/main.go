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
type Option struct {
	Index int    `json:"index"`
	Value string `json:"value"`
}

var ResponseMessages []ResponseMessage

func main() {

	ResponseMessages = []ResponseMessage{
		{
			Message: "1",
		},
		{
			Message: "2",
		},
	}
	http.HandleFunc("/", LandingPage)
	http.HandleFunc("/one", PrintLetterOne)
	http.HandleFunc("/two", PrintLetterTwo)
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
func PrintLetterTwo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responseMessage := ResponseMessages
		responseInByte, _ := json.Marshal(responseMessage)
		w.Write(responseInByte)
	} else if r.Method == http.MethodPost {
		if r.Body != nil {
			buffer := ResponseMessage{}
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&buffer)
			ResponseMessages = append(ResponseMessages, buffer)
			responseInByte, _ := json.Marshal(ResponseMessages)
			w.Write(responseInByte)
		}
	} else if r.Method == http.MethodDelete {
		if r.Body != nil {
			buffer := Option{}
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&buffer)
			for i := 0; i < len(ResponseMessages); i++ {
				if buffer.Index == i {
					ResponseMessages = append(ResponseMessages[:i], ResponseMessages[i:]...)
				}
			}
			responseInByte, _ := json.Marshal(ResponseMessages)
			w.Write(responseInByte)
		}
	} else if r.Method == http.MethodPut {
		if r.Body != nil {
			buffer := Option{}
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&buffer)
			ResponseMessages[buffer.Index].Message = buffer.Value
			responseInByte, _ := json.Marshal(ResponseMessages)
			w.Write(responseInByte)
		}
	}
}
