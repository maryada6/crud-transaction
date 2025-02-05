package handlers

import (
	"crud-transaction/db"
	"crud-transaction/helpers"
	"crud-transaction/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.ParseInt(r.URL.Path[len("/transactionservice/transaction/"):], 10, 64)

	var transaction models.Transaction
	if err := db.DB.First(&transaction, transactionID).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	typeName := r.URL.Path[len("/transactionservice/types/"):]

	var transactions []models.Transaction
	var transactionIDs []int64

	if err := db.DB.Where("type = ?", typeName).Find(&transactions).Error; err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}

	for _, transaction := range transactions {
		transactionIDs = append(transactionIDs, transaction.ID)
	}

	json.NewEncoder(w).Encode(transactionIDs)
}

func GetTransactionSum(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.ParseInt(r.URL.Path[len("/transactionservice/sum/"):], 10, 64)

	sum := helpers.CalculateSum(transactionID)

	json.NewEncoder(w).Encode(map[string]float64{"sum": sum})
}
