package environment

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func EnvVar(key string) string {
	env := os.Getenv("ENV")
	if env == "development" {
		viper.SetConfigFile(".development.env")
	} else {
		viper.SetConfigFile(".production.env")
	}

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if ok {
		return value
	} else {
		log.Printf("env var not found in .env file")
		log.Printf("looking at environment")
		systemEnv := os.Getenv(key)
		return systemEnv
	}

}

func IsDevelopment() bool {
	env := os.Getenv("ENV")
	if env == "development" {
		return true
	} else {
		return false
	}
}

func IsProduction() bool {
	return !IsDevelopment()
}
