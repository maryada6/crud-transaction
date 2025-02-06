package config

import (
	"log"

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

func init() {
	load()
}
