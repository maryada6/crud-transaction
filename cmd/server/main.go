package main

import (
	"fmt"
	"log"
	"net/http"

	"crud-transaction/config"
	routes "crud-transaction/router"
)

func main() {
	log.Println("Starting Transaction Service API")
	router := routes.SetupRouter()
	log.Println("Server is running on port", config.GetServerPort())
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", config.GetServerPort()), router))
}
