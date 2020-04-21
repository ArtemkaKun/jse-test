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
	decodeRequestJSON(req, &newMessage)

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
	decodeRequestJSON(req, &newMessage)

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
		sendError(writer, fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized))
		return
	}

	if !isIdExist(newMessage.Id) {
		sendError(writer, "User with the this ID doesn't exist!")
		return
	}

	deposit := createNewDeposit(newMessage)
	setNewBalance(deposit, newMessage)

	sendErrorWithBalance(writer, "", deposit.BalanceAfter)
}

func createNewDeposit(newMessage AddDepositMessage) Deposit {
	var newDeposit Deposit

	newDeposit.UserId = newMessage.Id
	newDeposit.DepositId = newMessage.DepositId
	newDeposit.BalanceBefore = GetUserBalance(newDeposit.UserId)
	newDeposit.BalanceAfter = newDeposit.BalanceBefore + newMessage.Amount
	newDeposit.DepositTime = time.Now()

	return newDeposit
}

func setNewBalance(newDeposit Deposit, newMessage AddDepositMessage) {
	SetUserBalance(newDeposit.UserId, newDeposit.BalanceAfter)
	IncreaseUserDepositCount(newDeposit.UserId)
	IncreaseUserDepositSum(newDeposit.UserId, newMessage.Amount)
	addNewDeposit(newDeposit)
}

func makeTransaction(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newMessage AddTransactionMessage
	err := json.NewDecoder(req.Body).Decode(&newMessage)
	DecodingJSONError(err)

	if newMessage.Token != TOKEN {
		sendError(writer, fmt.Sprintf("Autentification error: %v", http.StatusUnauthorized))
		return
	}

	if !isIdExist(newMessage.Id) {
		sendError(writer, "User with the this ID doesn't exist!")
		return
	}

	if GetUserBalance(newMessage.Id)-newMessage.Amount < 0 && newMessage.Type == "Bet" {
		sendError(writer, "Doesn't have enough money on balance to do this action!")
		return
	}

	transaction := createNewTransaction(newMessage)
	setNewTransaction(transaction, newMessage)

	sendErrorWithBalance(writer, "", transaction.BalanceAfter)
}

func setNewTransaction(newTransaction Transaction, newMessage AddTransactionMessage) {
	SetUserBalance(newTransaction.UserId, newTransaction.BalanceAfter)

	if newMessage.Type == "Bet" {
		IncreaseUserBetCount(newTransaction.UserId)
		IncreaseUserBetSum(newTransaction.UserId, newMessage.Amount)
		addNewTransaction(newTransaction)
		return
	}

	IncreaseUserWinCount(newTransaction.UserId)
	IncreaseUserWinSum(newTransaction.UserId, newMessage.Amount)
	addNewTransaction(newTransaction)
}

func createNewTransaction(newMessage AddTransactionMessage) Transaction {
	var newTransaction Transaction

	id := newMessage.Id
	amount := newMessage.Amount

	newTransaction.UserId = id
	newTransaction.TransactionId = newMessage.TransactionId
	newTransaction.Type = newMessage.Type
	newTransaction.Amount = amount
	newTransaction.BalanceBefore = GetUserBalance(id)

	if newMessage.Type == "Bet" {
		amount *= -1
	}

	newTransaction.BalanceAfter = newTransaction.BalanceBefore + amount
	newTransaction.TransactionTime = time.Now()

	return newTransaction
}

func decodeRequestJSON(req *http.Request, newMessage *NewUserMessage) {
	err := json.NewDecoder(req.Body).Decode(newMessage)
	DecodingJSONError(err)
}

func sendErrorWithBalance(writer http.ResponseWriter, errorText string, newBalance float32) {
	errorMessage := ErrorWithBalance{Error: errorText, Balance: newBalance}

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
