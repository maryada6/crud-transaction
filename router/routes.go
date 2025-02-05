package routes

import (
	"crud-transaction/db"
	"crud-transaction/handlers"

	"github.com/gorilla/mux"
)

const (
	GET = "GET"
	PUT = "PUT"
)

func SetupRouter() *mux.Router {
	db.InitDB()
	router := mux.NewRouter()

	router.Use(loggerMiddleware)
	router.HandleFunc("/transactionservice/transaction/{transaction_id}", handlers.CreateTransaction).Methods(PUT)
	router.HandleFunc("/transactionservice/transaction/{transaction_id}", handlers.GetTransaction).Methods(GET)
	router.HandleFunc("/transactionservice/types/{type}", handlers.GetTransactionsByType).Methods(GET)
	router.HandleFunc("/transactionservice/sum/{transaction_id}", handlers.GetTransactionSum).Methods(GET)

	return router
}
