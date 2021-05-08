package database

import (
	"fmt"
	"net/http"
	"trackly-backend/app/models"

	"trackly-backend/app/utils"

	"github.com/gorilla/mux"
)

func (db *DB) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)
	// db.First(&user, params["id"])
	if db_resp := db.First(&user, params["id"]); db_resp.Error != nil {
		// Create failed, do something e.g. return, panic etc.
		utils.ErrorResponse(w, 200, db_resp.Error.Error())

		fmt.Println(db_resp.Error)
		return
	}
	utils.JsonResponse(w, 200, user)
}

func (db *DB) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(users)
	utils.JsonResponse(w, 200, users)
}
