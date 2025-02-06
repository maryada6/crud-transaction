package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crud-transaction/config"
	"crud-transaction/db"
	"crud-transaction/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	config.Load()
	db.InitDB()

	t.Run("invalid request", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", nil)
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Request body is empty")
	})

	t.Run("invalid JSON", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", bytes.NewBuffer([]byte("invalid JSON")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid request payload")
	})

	t.Run("invalid transaction ID", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("POST", "/transactionservice/transaction/abc", nil)
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid transaction ID")
	})

	t.Run("invalid transaction type", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			Amount: 100,
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Transaction type is required")
	})

	t.Run("negative amount", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			Amount: -10,
			Type:   "shopping",
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/2", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Transaction amount must be greater than zero")
	})

	t.Run("duplicate transaction", func(t *testing.T) {
		defer truncateDB()
		db.DB.Create(&models.Transaction{ID: 3, Amount: 50, Type: "shopping"})
		transaction := models.Transaction{
			Amount: 50,
			Type:   "shopping",
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/3", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code)
		assert.Contains(t, resp.Body.String(), "Duplicate transaction")
	})

	t.Run("valid transaction with parent ID", func(t *testing.T) {
		defer truncateDB()
		db.DB.Create(&models.Transaction{ID: 3, Amount: 50, Type: "shopping"})
		transaction := models.Transaction{
			Type:     "shopping",
			Amount:   100.50,
			ParentID: int64(5),
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid parent transaction ID")
	})

	t.Run("valid transaction", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			Type:   "shopping",
			Amount: 100.50,
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]string
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])
	})

	t.Run("valid transaction with parent ID", func(t *testing.T) {
		defer truncateDB()
		db.DB.Create(&models.Transaction{ID: 3, Amount: 50, Type: "shopping"})
		transaction := models.Transaction{
			Type:     "shopping",
			Amount:   100.50,
			ParentID: int64(3),
		}
		jsonData, _ := json.Marshal(transaction)

		req, _ := http.NewRequest("POST", "/transactionservice/transaction/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		CreateTransaction(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]string
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])
	})
}

func truncateDB() {
	db.DB.Exec("TRUNCATE TABLE transactions")
}
