package config

import (

	"github.com/spf13/viper"
)

type Environment string

const (
	Development Environment = "development"
)

func LoadConfig(filePath string, fileName string) (err error) {
	viper.SetDefault("ENV", Development)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filePath)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return
	}
	return
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