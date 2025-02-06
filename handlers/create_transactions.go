package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-transaction/db"
	"crud-transaction/models"
)

func handleError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

func decodeRequestBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func isDuplicateTransaction(transactionID int64) bool {
	var existingTransaction models.Transaction
	err := db.DB.Where("id = ?", transactionID).First(&existingTransaction).Error
	return err == nil
}

func isValidParentTransaction(parentID int64) bool {
	var parentTransaction models.Transaction
	err := db.DB.Where("id = ?", parentID).First(&parentTransaction).Error
	return err == nil
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID, err := strconv.ParseInt(r.URL.Path[len("/transactionservice/transaction/"):], 10, 64)
	if err != nil {
		handleError(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		handleError(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var transaction models.Transaction
	if err := decodeRequestBody(r, &transaction); err != nil {
		handleError(w, "Invalid request payload", http.StatusUnprocessableEntity)
		return
	}

	if transaction.Type == "" {
		handleError(w, "Transaction type is required", http.StatusBadRequest)
		return
	}

	if transaction.Amount <= 0 {
		handleError(w, "Transaction amount must be greater than zero", http.StatusBadRequest)
		return
	}

	transaction.ID = transactionID
	if isDuplicateTransaction(transaction.ID) {
		handleError(w, "Duplicate transaction", http.StatusConflict)
		return
	}

	if transaction.ParentID != 0 && !isValidParentTransaction(transaction.ParentID) {
		handleError(w, "Invalid parent transaction ID", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		handleError(w, "Error saving transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		handleError(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
