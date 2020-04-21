package main

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

type NewUserMessage struct {
	Id      uint64  `json:"id"`
	Balance float32 `json:"balance"`
	Token   string  `json:"token"`
}

//type GetUserMessage struct {
//	Id    uint64 `json:"id"`
//	Token string `json:"token"`
//}

type AddDepositMessage struct {
	Id        uint64  `json:"userid"`
	DepositId uint64  `json:"depositId"`
	Amount    float32 `json:"amount"`
	Token     string  `json:"token"`
}

type Deposit struct {
	UserId        uint64
	DepositId     uint64
	BalanceBefore float32
	BalanceAfter  float32
	DepositTime   string
}

type AddReqError struct {
	Error   string  `json:"error"`
	Balance float32 `json:"balance"`
}
