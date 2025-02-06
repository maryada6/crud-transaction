package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-transaction/db"
	"crud-transaction/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID, err := strconv.ParseInt(r.URL.Path[len("/transactionservice/transaction/"):], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var transaction models.Transaction
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		http.Error(w, "Transaction amount must be greater than zero", http.StatusBadRequest)
		return
	}

	transaction.ID = transactionID
	var existingTransaction models.Transaction
	if err := db.DB.Where("id = ?", transaction.ID).First(&existingTransaction).Error; err == nil {
		http.Error(w, "Duplicate transaction", http.StatusConflict)
		return
	}

	if transaction.ParentID != nil {
		var parentTransaction models.Transaction
		if err := db.DB.Where("id = ?", *transaction.ParentID).First(&parentTransaction).Error; err != nil {
			http.Error(w, "Invalid parent transaction ID", http.StatusBadRequest)
			return
		}
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		http.Error(w, "Error saving transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
