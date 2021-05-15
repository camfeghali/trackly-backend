package main

import (
	"fmt"
	"trackly-backend/app/utils"

	"github.com/spf13/viper"
)

type Configuration struct {
	ENV           string
	PORT          int
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_NAME       string
	AUTHORIZATION bool
}

type Envconfig struct {
	PORT          int
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_NAME       string
	AUTHORIZATION bool
	// DB_PORT       string
	// DB_HOST       string
}

func GetConfig() {
	vi := getConfigfile("config.json", "./app/main")
	err := vi.ReadInConfig()
	utils.CheckError(err)

	configuration := Configuration{}
	vi.Unmarshal(&configuration)

	fmt.Println(configuration)
}

func getConfigfile(fileName, directory string) *viper.Viper {
	vi := viper.New()
	vi.SetConfigType("json")
	vi.AddConfigPath("./app/main")
	vi.SetConfigName("config")
	return vi
}

// {
//     "dev": {
//         "port": 8080,
//         "db_user": "trackly_user",
//         "db_pw": "insabgho123",
//         "db_name": "trackly",
//         "authorization": false
//     }
// }
