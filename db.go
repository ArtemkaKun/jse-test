package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

var connection *pgx.Conn

func init() {
	var err error
	connection, err = pgx.Connect(context.Background(), "postgres://postgres:1337@localhost:5432/jse-test")
	if err != nil {
		log.Panic(fmt.Errorf("Unable to connection to database: %v\n", err))
	} else {
		fmt.Println("Connected to PSQL!")
	}
}

func addNewUserToDB(newUser User) {
	_, err := connection.Exec(context.Background(), "INSERT INTO users VALUES($1, $2)", newUser.Id, newUser.Balance)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}

func addNewDeposit(newDeposit Deposit) {
	_, err := connection.Exec(context.Background(), "INSERT INTO deposits VALUES($1, $2, $3, $4, $5)", newDeposit.DepositId, newDeposit.UserId,
		newDeposit.BalanceBefore, newDeposit.BalanceAfter, newDeposit.DepositTime)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}

func addNewTransaction(newTransaction Transaction) {
	_, err := connection.Exec(context.Background(), "INSERT INTO transactions VALUES($1, $2, $3, $4, $5, $6, $7)", newTransaction.TransactionId, newTransaction.UserId,
		newTransaction.Type, newTransaction.Amount, newTransaction.BalanceBefore, newTransaction.BalanceAfter, newTransaction.TransactionTime)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}

func updateUser(newUser *User) {
	_, err := connection.Exec(context.Background(), "UPDATE users SET balance = $1 WHERE id = $2", newUser.Balance, newUser.Id)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}
