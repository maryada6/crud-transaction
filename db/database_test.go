package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	t.Run("should connect to database", func(t *testing.T) {
		db = nil
		GetDB()
		assert.NotNil(t, db)
	})
}
