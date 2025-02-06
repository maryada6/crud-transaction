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
	port := config.GetServerPort()
	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
