package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-transaction/db"
	"crud-transaction/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.ParseInt(r.URL.Path[len("/transactionservice/transaction/"):], 10, 64)

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	transaction.ID = transactionID

	var existingTransaction models.Transaction
	if err := db.DB.Where("id = ?", transaction.ID).First(&existingTransaction).Error; err == nil {
		http.Error(w, "Duplicate transaction", http.StatusConflict)
		return
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		http.Error(w, "Error saving transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
