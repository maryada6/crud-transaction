package routes

import (
	"crud-transaction/handlers"

	"github.com/gorilla/mux"
)

const (
	GET = "GET"
	PUT = "PUT"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggerMiddleware)
	trnHandler := handlers.NewTransactionHandler()

	router.HandleFunc("/transactionservice/transaction/{transaction_id}", trnHandler.CreateTransaction).Methods(PUT)
	router.HandleFunc("/transactionservice/transaction/{transaction_id}", trnHandler.GetTransaction).Methods(GET)
	router.HandleFunc("/transactionservice/types/{type}", trnHandler.GetTransactionsByType).Methods(GET)
	router.HandleFunc("/transactionservice/sum/{transaction_id}", trnHandler.GetTransactionSum).Methods(GET)

	return router
}
