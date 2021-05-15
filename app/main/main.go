package main

import (
	"fmt"
	"trackly-backend/app/datastore"
	"trackly-backend/app/security"
	"trackly-backend/app/utils"
)

func main() {

	config := GetConfig("dev")
	if config.AUTHORIZATION_ENABLED {
		fmt.Println("JWT Authorization is enabled, you can use this token to make authorized requests:")
		JWTToken, err := security.GenerateJWT()
		utils.CheckError(err)
		fmt.Println(JWTToken)
	} else {
		fmt.Println("JWT Authorization is disabled")
	}

	db := datastore.NewDBInstance(config.DB_USERNAME, config.DB_PASSWORD, config.DB_ADDRESS, config.DB_NAME, config.DB_PORT)
	err := db.RunMigrations()
	utils.CheckError(err)
	handleRequests(config.PORT, db)
}
