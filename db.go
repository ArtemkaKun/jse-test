package jse_test

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

var connection *pgx.Conn

func init() {
	var err error
	connection, err = pgx.Connect(context.Background(), "postgres://postgres:1337@localhost:5432/jse_test")
	if err != nil {
		log.Panic(fmt.Errorf("Unable to connection to database: %v\n", err))
	} else {
		fmt.Println("Connected to PSQL!")
	}
}

func addNewUserToDB(new_user User) {
	_, err := connection.Exec(context.Background(), "INSERT INTO users VALUES($1, $2)", new_user.Id, new_user.Balance)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while inserting: %v\n", err))
	}
}
