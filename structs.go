package jse_test

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

type GetUserMessage struct {
	Id    uint64 `json:"id"`
	Token string `json:"token"`
}

type ReqError struct {
	Error string `json:"error"`
}
