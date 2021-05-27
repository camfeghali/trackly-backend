package datastore

import (
	"net/http"
	"time"
	"trackly-backend/app/security"
	"trackly-backend/app/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Password  string         `json:"-" gorm:"not null"`
	Clients   []*Client      `json:"clients" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
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
	password := r.FormValue("password")

	if password != "" {
		encryptedPassword, err := security.Encrypt(password)
		if err != nil {
			utils.ErrorResponse(w, 200, err.Error())
			return
		}
		user := User{FirstName: firstName, LastName: lastName, Password: encryptedPassword}
		if db_resp := db.Create(&user); db_resp.Error != nil {
			utils.ErrorResponse(w, 200, db_resp.Error.Error())
			return
		}
		utils.JsonResponse(w, 200, user)
	} else {
		utils.ErrorResponse(w, 200, "Password not provided")
	}
}

func (db *DB) Login(w http.ResponseWriter, r *http.Request) {
	var foundUser User
	firstName := r.FormValue("firstName")
	password := r.FormValue("password")

	if db_resp := db.Where("first_name = ?", firstName).First(&foundUser); db_resp.Error != nil {
		utils.ErrorResponse(w, 200, db_resp.Error.Error())
		return
	} else {
		if ok, err := (security.PasswordMatches(password, foundUser.Password)); ok {
			// Send token as well!
			utils.JsonResponse(w, 200, foundUser)
		} else {
			utils.ErrorResponse(w, 200, err.Error())
		}
	}
}
