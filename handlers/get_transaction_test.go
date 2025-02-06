package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crud-transaction/config"
	"crud-transaction/db"
	"crud-transaction/models"

	"github.com/stretchr/testify/assert"
)

func TestGetTransaction(t *testing.T) {
	config.Load()
	db.InitDB()

	transactionHandler := NewTransactionHandler()

	t.Run("valid transaction", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			ID:     1,
			Type:   "shopping",
			Amount: 100.50,
		}
		db.DB.Create(&transaction)

		req, _ := http.NewRequest("GET", "/transactionservice/transaction/1", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransaction(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response models.Transaction
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, transaction.ID, response.ID)
		assert.Equal(t, transaction.Type, response.Type)
		assert.Equal(t, transaction.Amount, response.Amount)
	})

	t.Run("invalid transaction ID", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/transaction/abc", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransaction(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid transaction ID")
	})

	t.Run("transaction not found", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/transaction/999", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransaction(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}

func TestGetTransactionsByType(t *testing.T) {
	config.Load()
	db.InitDB()

	transactionHandler := NewTransactionHandler()
	t.Run("transactions by type", func(t *testing.T) {
		defer truncateDB()
		db.DB.Create(&models.Transaction{ID: 1, Type: "shopping", Amount: 100.50})
		db.DB.Create(&models.Transaction{ID: 2, Type: "shopping", Amount: 50.75})
		db.DB.Create(&models.Transaction{ID: 3, Type: "food", Amount: 30.20})

		req, _ := http.NewRequest("GET", "/transactionservice/types/shopping", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionsByType(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response []int64
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Contains(t, response, int64(1))
		assert.Contains(t, response, int64(2))
	})

	t.Run("empty transaction type", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/types/", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionsByType(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Transaction type is required")
	})

	t.Run("no transactions for type", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/types/nonexistent", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionsByType(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response []int64
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 0)
	})
}

func TestGetTransactionSum(t *testing.T) {
	config.Load()
	db.InitDB()

	transactionHandler := NewTransactionHandler()
	t.Run("valid transaction sum", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			ID:     1,
			Type:   "shopping",
			Amount: 100.50,
		}
		db.DB.Create(&transaction)

		req, _ := http.NewRequest("GET", "/transactionservice/sum/1", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionSum(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]float64
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 100.50, response["sum"])
	})

	t.Run("valid transaction sum with multiple transactions", func(t *testing.T) {
		defer truncateDB()
		db.DB.Create(&models.Transaction{ID: 1, Type: "shopping", Amount: 100.50})
		db.DB.Create(&models.Transaction{ID: 2, Type: "shopping", Amount: 50.75, ParentID: 1})
		db.DB.Create(&models.Transaction{ID: 3, Type: "food", Amount: 30.20, ParentID: 1})
		db.DB.Create(&models.Transaction{ID: 4, Type: "food", Amount: 20.15, ParentID: 2})
		db.DB.Create(&models.Transaction{ID: 5, Type: "food", Amount: 10.10, ParentID: 3})

		req, _ := http.NewRequest("GET", "/transactionservice/sum/1", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionSum(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]float64
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 211.7, response["sum"])
	})

	t.Run("invalid transaction ID", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/sum/abc", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionSum(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid transaction ID")
	})

	t.Run("transaction sum not found", func(t *testing.T) {
		defer truncateDB()
		req, _ := http.NewRequest("GET", "/transactionservice/sum/999", nil)
		resp := httptest.NewRecorder()

		transactionHandler.GetTransactionSum(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]float64
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, float64(0), response["sum"])
	})
}
