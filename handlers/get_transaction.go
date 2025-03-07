package handlers

import (
	"crud-transaction/db"
	"crud-transaction/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func (t *transactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID, err := strconv.ParseInt(r.URL.Path[len("/transactionservice/transaction/"):], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	if transactionID <= 0 {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var transaction models.Transaction
	if err := db.GetDB().First(&transaction, transactionID).Error; err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (t *transactionHandler) GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	typeName := r.URL.Path[len("/transactionservice/types/"):]

	if typeName == "" {
		http.Error(w, "Transaction type is required", http.StatusBadRequest)
		return
	}

	var transactionIDs []int64
	if err := db.GetDB().Model(&models.Transaction{}).
		Where("type = ?", typeName).
		Pluck("id", &transactionIDs).Error; err != nil {
		http.Error(w, "Error retrieving transaction IDs", http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(transactionIDs)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (t *transactionHandler) GetTransactionSum(w http.ResponseWriter, r *http.Request) {
	transactionID, err := strconv.ParseInt(r.URL.Path[len("/transactionservice/sum/"):], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	if transactionID <= 0 {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	sum := calculateSum(transactionID)

	err = json.NewEncoder(w).Encode(map[string]float64{"sum": sum})
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
