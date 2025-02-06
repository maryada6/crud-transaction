package routes

import (
	"crud-transaction/db"
	"crud-transaction/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	t.Run("valid GET transaction", func(t *testing.T) {
		defer truncateDB()
		transaction := models.Transaction{
			ID:     1,
			Type:   "shopping",
			Amount: 100.50,
		}
		db.GetDB().Create(&transaction)
		req, _ := http.NewRequest("GET", "/transactionservice/transaction/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
		var response models.Transaction
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, transaction.ID, response.ID)
		assert.Equal(t, transaction.Type, response.Type)
		assert.Equal(t, transaction.Amount, response.Amount)
	})

	t.Run("invalid route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/transactionservice/nonexistent", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}

func truncateDB() {
	db.GetDB().Exec("TRUNCATE TABLE transactions")
}
