package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func load() {
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	configFile := "application"
	env := GetStringWithDefault("ENV", "dev")

	if env == "test" {
		configFile = "application_test"
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType("yml")

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func GetServerPort() int {
	return GetIntWithDefault("PORT", 0)
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

func init() {
	load()
}
