package main

import (
	"trackly-backend/app/database"
)

func main() {
	var port string = "8080"
	db := database.NewDBInstance("trackly_user", "insabgho123", "trackly")
	handleRequests(port, db)
}
