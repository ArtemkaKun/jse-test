package main

import "time"

type User struct {
	Id      uint64  `json:"id"`
	Balance float32 `json:"balance"`
}

type UserStats struct {
	DepositCount uint32  `json:"depositCount"`
	DepositSum   float32 `json:"depositSum"`
	BetCount     uint32  `json:"betCount"`
	BetSum       float32 `json:"betSum"`
	WinCount     uint32  `json:"winCount"`
	WinSum       float32 `json:"winSum"`
}

type AllUserStats struct {
	Id           uint64  `json:"id"`
	Balance      float32 `json:"balance"`
	DepositCount uint32  `json:"depositCount"`
	DepositSum   float32 `json:"depositSum"`
	BetCount     uint32  `json:"betCount"`
	BetSum       float32 `json:"betSum"`
	WinCount     uint32  `json:"winCount"`
	WinSum       float32 `json:"winSum"`
}

type Deposit struct {
	UserId        uint64
	DepositId     uint64
	BalanceBefore float32
	BalanceAfter  float32
	DepositTime   time.Time
}

type Transaction struct {
	UserId          uint64
	TransactionId   uint64
	Type            string
	Amount          float32
	BalanceBefore   float32
	BalanceAfter    float32
	TransactionTime time.Time
}

type NewUserMessage struct {
	Id      uint64  `json:"id"`
	Balance float32 `json:"balance"`
	Token   string  `json:"token"`
}

type AddDepositMessage struct {
	Id        uint64  `json:"userId"`
	DepositId uint64  `json:"depositId"`
	Amount    float32 `json:"amount"`
	Token     string  `json:"token"`
}

type AddTransactionMessage struct {
	Id            uint64  `json:"userId"`
	TransactionId uint64  `json:"transactionId"`
	Type          string  `json:"type"`
	Amount        float32 `json:"amount"`
	Token         string  `json:"token"`
}

type ErrorWithBalance struct {
	Error   string  `json:"error"`
	Balance float32 `json:"balance"`
}
