package main

import (
	"trackly-backend/app/datastore"
	"trackly-backend/app/utils"
)

func main() {

	GetConfig()
	var port string = "8080"
	db := datastore.NewDBInstance("trackly_user", "insabgho123", "trackly")
	err := db.RunMigrations()
	utils.CheckError(err)
	handleRequests(port, db)
}
