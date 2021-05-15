package main

import (
	"trackly-backend/app/datastore"
	"trackly-backend/app/utils"
)

func main() {

	config := GetConfig("dev")

	db := datastore.NewDBInstance(config.DB_USERNAME, config.DB_PASSWORD, config.DB_ADDRESS, config.DB_NAME, config.DB_PORT)
	err := db.RunMigrations()
	utils.CheckError(err)
	handleRequests(config.PORT, db)
}
