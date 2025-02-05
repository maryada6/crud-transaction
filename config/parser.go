package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func GetStringOrPanic(key string) string {
	val := viper.GetString(key)
	if val == "" {
		panic(fmt.Errorf("key %s is mandatory", key))
	}

	return val
}

func GetStringWithDefault(key string, defaultValue string) string {
	val := viper.GetString(key)
	if val == "" {
		val = defaultValue
	}

	return val
}

func GetIntWithDefault(key string, defaultValue int) int {
	v := viper.GetString(key)
	value, err := strconv.Atoi(v)
	if err != nil {
		value = defaultValue
	}

	return value
}
