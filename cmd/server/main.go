package main

import (
	"fmt"
	"log"
	"net/http"

	"crud-transaction/config"
	"crud-transaction/db"
	routes "crud-transaction/router"
)

func main() {
	config.Load()
	db.InitDB()

	log.Println("Starting Transaction Service API...")
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", config.GetServerPort()), router))
}
