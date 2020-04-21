package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var router *mux.Router
var TOKEN = "testtask"

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
	decodeRequestJSON(req, newMessage)

	if newMessage.Token != TOKEN {
		sendError(writer, fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized))
		return
	}

	if isIdExist(newMessage.Id) {
		sendError(writer, "User with the same ID already exist!")
		return
	}

	newUser := User{Id: newMessage.Id, Balance: newMessage.Balance}

	addNewUser(newUser)

	sendError(writer, "")
}

func addNewUser(newUser User) {
	addNewUserToBuffer(newUser)
	addNewUserToDB(newUser)
}

func getUserInfo(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newMessage NewUserMessage
	decodeRequestJSON(req, newMessage)

	if newMessage.Token != TOKEN {
		sendError(writer, fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized))
		return
	}

	if !isIdExist(newMessage.Id) {
		sendError(writer, "User with the same ID doesn't exist!")
		return
	}

	allStats := GetAllUserInfo(newMessage.Id)

	sendAllStats(writer, allStats)
}

func sendAllStats(writer http.ResponseWriter, allStats AllUserStats) {
	err := json.NewEncoder(writer).Encode(allStats)
	if err != nil {
		EncodingJSONError(err)
	}
}

func addUserDeposit(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newMessage AddDepositMessage
	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)

	if newMessage.Token != TOKEN {
		sendAddError(writer, fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized), 0)
		return
	}

	if !isIdExist(newMessage.Id) {
		sendAddError(writer, "User with the this ID doesn't exist!", 0)
		return
	}

	deposit := createNewDeposit(newMessage)
	setNewBalance(deposit, newMessage)

	sendAddError(writer, "", deposit.BalanceAfter)
}

func createNewDeposit(newMessage AddDepositMessage) Deposit {
	var newDeposit Deposit

	newDeposit.UserId = newMessage.Id
	newDeposit.DepositId = newMessage.DepositId
	newDeposit.BalanceBefore = GetUserBalance(newDeposit.UserId)
	newDeposit.BalanceAfter = newDeposit.BalanceBefore + newMessage.Amount
	newDeposit.DepositTime = fmt.Sprintf("%v", time.Now())

	return newDeposit
}

func setNewBalance(newDeposit Deposit, newMessage AddDepositMessage) {
	SetUserBalance(newDeposit.UserId, newDeposit.BalanceAfter)
	IncreaseUserDepositCount(newDeposit.UserId)
	IncreaseUserDepositAmount(newDeposit.UserId, newMessage.Amount)
}

func makeTransaction(writer http.ResponseWriter, req *http.Request) {

}

func decodeRequestJSON(req *http.Request, newMessage NewUserMessage) {
	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)
}

func sendAddError(writer http.ResponseWriter, errorText string, newBalance float32) {
	errorMessage := AddReqError{Error: errorText, Balance: newBalance}

	err := json.NewEncoder(writer).Encode(errorMessage)
	if err != nil {
		EncodingJSONError(err)
	}
}

func sendError(writer http.ResponseWriter, errorText string) {
	errorMessage := map[string]string{"error": errorText}

	err := json.NewEncoder(writer).Encode(errorMessage)
	if err != nil {
		EncodingJSONError(err)
	}
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}

func DecodingJSONError(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
	}
}
