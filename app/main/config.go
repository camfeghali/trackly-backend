package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"trackly-backend/app/utils"
)

type AppConfiguration struct {
	ENV DevConfiguration
}

type DevConfiguration struct {
	DEV     Configuration
	STAGING Configuration
	PROD    Configuration
}

type Configuration struct {
	PORT                  int    `json:"port"`
	DB_USERNAME           string `json:"db_username"`
	DB_PASSWORD           string `json:"db_password"`
	DB_NAME               string `json:"db_name"`
	DB_ADDRESS            string `json:"db_address"`
	DB_PORT               int    `json:"db_port"`
	AUTHORIZATION_ENABLED bool   `json:"authorization_enabled"`
}

func LoadConfig() AppConfiguration {
	var configuration AppConfiguration
	jsonFile, err := os.Open("./app/main/config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	utils.CheckError(err)
	json.Unmarshal([]byte(byteValue), &configuration)
	return configuration
}

func GetConfig(env string) Configuration {
	appConfig := LoadConfig()
	switch env {
	case "dev":
		return appConfig.ENV.DEV
	case "prod":
		return appConfig.ENV.PROD
	case "staging":
		return appConfig.ENV.STAGING
	default:
		return appConfig.ENV.DEV
	}
}
