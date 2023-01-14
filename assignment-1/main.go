package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	_ "Hacktiv8project/assignment-1/docs"
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
	r := mux.NewRouter()
	ResponseMessages = []ResponseMessage{
		{
			Message: "1",
		},
		{
			Message: "2",
		},
	}
	r.HandleFunc("/", LandingPage).Methods(http.MethodGet)
	r.HandleFunc("/messages/Get", GorillaGet).Methods(http.MethodGet)
	r.HandleFunc("/messages/{id}", GorillaGetById).Methods(http.MethodGet)
	r.HandleFunc("/messages/Post", GorillaPost).Methods(http.MethodPost)
	r.HandleFunc("/messages/{id}", GorillaPut).Methods(http.MethodPut)
	r.HandleFunc("/messages/{id}", GorillaDelete).Methods(http.MethodDelete)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	http.HandleFunc("/one", PrintLetterOne)
	http.HandleFunc("/two", PrintLetterTwo)
	fmt.Println("service started")
	log.Fatal(http.ListenAndServe(Address, r))
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

// Get is a handler for get messages API
// @Summary Get new messages
// @Description get all message list
// @Tags Basic CRUD
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /messages [get]
func GorillaGet(w http.ResponseWriter, r *http.Request) {
	responseMessage := ResponseMessages
	responseInByte, _ := json.Marshal(responseMessage)
	w.Write(responseInByte)
}

// GetByID is a handler for get messages API
// @Summary GetByID new messages
// @Description get all message list
// @Tags Basic CRUD
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /messages/{id} [get]
func GorillaGetById(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["id"])

	responseMessage := ResponseMessages[index]
	responseInByte, _ := json.Marshal(responseMessage)
	w.Write(responseInByte)
}

// Create is a handler for create messages API
// @Summary Create new messages
// @Description get string by ID
// @Tags Basic CRUD
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /messages [post]
func GorillaPost(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		buffer := ResponseMessage{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&buffer)
		ResponseMessages = append(ResponseMessages, buffer)
		message := "Succesfully add message: " + buffer.Message
		w.Write([]byte(message))
	}
}

// Put is a handler for create messages API
// @Summary Update new messages
// @Description Update messages by ID
// @Tags Basic CRUD
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /messages/{id} [put]
func GorillaPut(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		buffer := ResponseMessage{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&buffer)
		index, _ := strconv.Atoi(mux.Vars(r)["id"])
		ResponseMessages[index].Message = buffer.Message
		message := "Succesfully edit message: " + buffer.Message + " in index :" + strconv.Itoa(index)
		w.Write([]byte(message))
	}
}

// Delete is a handler for create messages API
// @Summary Delete new messages
// @Description Delete string by ID
// @Tags Basic CRUD
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /messages/{id} [delete]
func GorillaDelete(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		buffer := ResponseMessage{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&buffer)
		index, _ := strconv.Atoi(mux.Vars(r)["id"])
		for i := 0; i < len(ResponseMessages); i++ {
			if index == i {
				ResponseMessages = append(ResponseMessages[:i], ResponseMessages[i+1:]...)
			}
		}
		message := "Succesfully delete message in index:" + strconv.Itoa(index)
		w.Write([]byte(message))
	}
}
