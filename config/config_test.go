package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerPort(t *testing.T) {
	t.Run("should load test config", func(t *testing.T) {
		resetFn := SetAndLoad("ENV", "test")
		defer resetFn()
		assert.Equal(t, 0, GetServerPort())
		assert.Equal(t, "127.0.0.1", GetDatabaseHost())
		assert.Equal(t, 5432, GetDatabasePort())
		assert.Equal(t, "postgres", GetDatabaseUser())
		assert.Equal(t, "password", GetDatabasePassword())
		assert.Equal(t, "transaction_db_test", GetDatabaseName())
	})
}

func TestGetStringOrPanic(t *testing.T) {
	t.Run("should panic", func(t *testing.T) {
		assert.Panics(t, func() { GetStringOrPanic("TEST") })
	})

	t.Run("should not panic", func(t *testing.T) {
		resetFn := SetAndLoad("ENV", "test")
		defer resetFn()
		assert.NotPanics(t, func() { GetStringOrPanic("ENV") })
	})
}

func SetAndLoad(key, value string) func() {
	resetFn := UnsetAndLoad(key)

	os.Setenv(key, value)
	load()

	return resetFn
}

func UnsetAndLoad(key string) func() {
	originalValue := os.Getenv(key)
	os.Unsetenv(key)
	load()

	return func() {
		os.Setenv(key, originalValue)
		load()
	}
}
