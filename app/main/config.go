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

// func GetConfig() Configuration {
// 	vi := getConfigfile("config.json", "./app/main")
// 	err := vi.ReadInConfig()
// 	utils.CheckError(err)

// 	configuration := Configuration{}

// 	err = vi.Unmarshal(&configuration)
// 	utils.CheckError(err)

// 	return configuration
// }

// func getConfigfile(fileName, directory string) *viper.Viper {
// 	vi := viper.New()
// 	vi.SetConfigType("json")
// 	vi.AddConfigPath("./app/main")
// 	vi.SetConfigName("config")
// 	return vi
// }

// {
//     "species": "pigeon",
//     "decription": "likes to perch on rocks",
//     "dimensions": {
//         "height": 24,
//         "width": 10
//     }
// }

// {
//     "env": {
//         "dev": {
//             "port": 8080,
//             "db_user": "trackly_user",
//             "db_pw": "insabgho123",
//             "db_name": "trackly",
//             "authorization_enabled": false
//         }
//     }
// }
