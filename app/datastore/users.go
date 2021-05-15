package datastore

import (
	"net/http"
	"trackly-backend/app/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Clients   []*Client `gorm:"foreignKey:UserID"`
	gorm.Model
}

func (db *DB) GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	params := mux.Vars(r)
	if db_resp := db.First(&user, params["id"]); db_resp.Error != nil {
		utils.ErrorResponse(w, 200, db_resp.Error.Error())
		return
	}
	utils.JsonResponse(w, 200, user)
}

func (db *DB) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	if db_resp := db.Find(&users); db_resp.Error != nil {
		utils.ErrorResponse(w, 200, db_resp.Error.Error())
		return
	}
	utils.JsonResponse(w, 200, users)
}

func (db *DB) CreateUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")

	user := User{FirstName: firstName, LastName: lastName}

	if db_resp := db.Create(&user); db_resp.Error != nil {
		utils.ErrorResponse(w, 200, db_resp.Error.Error())
		return
	}

	utils.JsonResponse(w, 200, user)
}
