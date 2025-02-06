package routes

import (
	"crud-transaction/config"
	"crud-transaction/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	config.Load()
	db.InitDB()
	router := SetupRouter()

	t.Run("valid GET transaction", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/transactionservice/transaction/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("invalid route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/transactionservice/nonexistent", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
