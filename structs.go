package jse_test

type User struct {
	Id      uint64  `json:"id"`
	Balance float32 `json:"balance"`
}

type NewUserMessage struct {
	Id      uint64  `json:"id"`
	Balance float32 `json:"balance"`
	Token   string  `json:"token"`
}

type ReqError struct {
	Error string `json:"error"`
}
