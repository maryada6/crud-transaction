package db

import (
	"crud-transaction/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	t.Run("should connect to database", func(t *testing.T) {
		db = nil
		GetDB()
		assert.NotNil(t, db)
	})
	t.Run("should panic if cannot connect to database", func(t *testing.T) {
		db = nil
		config.SetAndLoad("DATABASE_USER", "invalid")
		assert.Panics(t, func() { GetDB() })
	})
}
