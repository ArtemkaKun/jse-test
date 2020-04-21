package jse_test

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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

func addUserDeposit(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newMessage AddDepositMessage

	errorMessage := AddReqError{Error: "", Balance: 0}

	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)

	if newMessage.Token != "testtask" {
		errorMessage = AddReqError{Error: fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized)}
	}

	if !isIdExist(newMessage.Id) {
		errorMessage = AddReqError{Error: "User with the this ID doesn't exist!"}
	}

	if errorMessage.Error != "" {
		sendAddError(writer, errorMessage)
		return
	}

	var newDeposit Deposit
	newDeposit.UserId = newMessage.Id
	newDeposit.DepositId = newMessage.DepositId
	newDeposit.BalanceBefore = GetUserBalance(newDeposit.UserId)
	newDeposit.BalanceAfter = newDeposit.BalanceBefore + newMessage.Amount
	newDeposit.DepositTime = fmt.Sprintf("%v", time.Now())

	SetUserBalance(newDeposit.UserId, newDeposit.BalanceAfter)
	IncreaseUserDepositCount(newDeposit.UserId)
	IncreaseUserDepositAmount(newDeposit.UserId, newMessage.Amount)

	errorMessage.Balance = newDeposit.BalanceAfter
	sendAddError(writer, errorMessage)
}

func DecodingJSONError(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
	}
}

func sendAddError(writer http.ResponseWriter, errorMessage AddReqError) {
	err := json.NewEncoder(writer).Encode(errorMessage)
	if err != nil {
		EncodingJSONError(err)
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
