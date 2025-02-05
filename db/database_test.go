package db

import (
	"crud-transaction/config"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	t.Run("should connect to database", func(t *testing.T) {
		DB = nil
		config.Load()
		InitDB()
	})
	t.Run("should panic if cannot connect to database", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		DB = nil
		setAndLoad("DATABASE_USER", "invalid")
		InitDB()
	})
}

func setAndLoad(key, value string) func() {
	resetFn := unsetAndLoad(key)

	os.Setenv(key, value)
	config.Load()

	return resetFn
}

func unsetAndLoad(key string) func() {
	originalValue := os.Getenv(key)
	os.Unsetenv(key)
	config.Load()

	return func() {
		os.Setenv(key, originalValue)
		config.Load()
	}
}
