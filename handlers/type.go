package handlers

import "net/http"

type TransactionInterface interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionsByType(w http.ResponseWriter, r *http.Request)
	GetTransactionSum(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct{}

var trnHandler *transactionHandler

func NewTransactionHandler() TransactionInterface {
	if trnHandler == nil {
		trnHandler = &transactionHandler{}
	}
	return trnHandler
}
