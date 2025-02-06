package db

import (
	"crud-transaction/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	t.Run("should connect to database", func(t *testing.T) {
		DB = nil
		InitDB()
	})
	t.Run("should panic if cannot connect to database", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		DB = nil
		config.SetAndLoad("DATABASE_USER", "invalid")
		InitDB()
	})
}
