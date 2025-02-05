package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	var configFile string
	env := GetStringWithDefault("ENV", "test")
	if env == "test" {
		configFile = "application_test"
	} else {
		configFile = "application"
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
	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

func GetServerPort() int {
	return GetIntWithDefault("PORT", 3000)
}
