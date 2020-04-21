package main

import (
	"log"
	"net/http"
)

func main() {
	go CheckChanges()
	log.Fatal(http.ListenAndServe(":8000", router))
}
