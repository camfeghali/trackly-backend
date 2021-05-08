package database

import (
	"encoding/json"
	"net/http"
	"trackly-backend/app/models"

	"github.com/gorilla/mux"
)

func (db *DB) GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)
	db.First(&user, params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (db *DB) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
