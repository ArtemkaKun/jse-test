package jse_test

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
	router.HandleFunc("/user/create", createNewUser).Methods("POST")
	router.HandleFunc("/user/get", getUserInfo).Methods("POST")
	router.HandleFunc("/user/deposit", addUserDeposit).Methods("POST")
	router.HandleFunc("/transaction", makeTransaction).Methods("POST")
}

func createNewUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newMessage NewUserMessage
	var newUser User

	errorMessage := ReqError{Error: ""}

	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)

	if newMessage.Token != "testtask" {
		errorMessage = ReqError{Error: fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized)}
	}

	if isIdExist(newMessage.Id) {
		errorMessage = ReqError{Error: "User with the same ID already exist!"}
	}

	if errorMessage.Error != "" {
		sendError(writer, errorMessage)
		return
	}

	newUser = User{Id: newMessage.Id, Balance: newMessage.Balance}

	addNewUserToBuffer(newUser)
	addNewUserToDB(newUser)

	sendError(writer, errorMessage)
}

func getUserInfo(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newMessage GetUserMessage

	errorMessage := ReqError{Error: ""}

	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)

	if newMessage.Token != "testtask" {
		errorMessage = ReqError{Error: fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized)}
	}

	if !isIdExist(newMessage.Id) {
		errorMessage = ReqError{Error: "User with the this ID doesn't exist!"}
	}

	if errorMessage.Error != "" {
		sendError(writer, errorMessage)
		return
	}

	all_stats := GetAllUserInfo(newMessage.Id)

	err = json.NewEncoder(writer).Encode(all_stats)
	if err != nil {
		EncodingJSONError(err)
	}
}

func DecodingJSONError(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
	}
}

func sendError(writer http.ResponseWriter, errorMessage ReqError) {
	err := json.NewEncoder(writer).Encode(errorMessage)
	if err != nil {
		EncodingJSONError(err)
	}
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}
