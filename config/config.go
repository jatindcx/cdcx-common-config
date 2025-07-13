package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, skipping...")
	}

	viper.AutomaticEnv()

	viper.SetDefault("APP_NAME", "DefaultApp")
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("DEBUG", false)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}