package main

import (
	"log"
	"net/http"

	"crud-transaction/config"
	routes "crud-transaction/router"
)

func main() {
	config.Load()
	log.Println("Starting Transaction Service API...")
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
