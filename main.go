package main

import (
	"log"
	"net/http"

	routes "crud-transaction/router"
)

func main() {
	log.Println("Starting Transaction Service API...")
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
