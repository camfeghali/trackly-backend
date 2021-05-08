package main

import (
	"fmt"
	"log"
	"net/http"
	"trackly-backend/app/database"
	"trackly-backend/app/models"
)

func main() {
	var port string = "8080"
	db := database.NewDBInstance("trackly_user", "insabgho123", "trackly")
	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	db.Create(&models.User{FirstName: "D42", LastName: "Doodoo"})

	// Read
	var user models.User
	// db.First(&user)                            // find user with integer primary key
	// db.First(&user, "last_name = ?", "Doodoo") // find user with code D42

	// // Update - update user's price to 200
	// db.Model(&user).Update("last_name", "Kathy")
	// // Update - update multiple fields
	// db.Model(&user).Updates(User{FirstName: "Emmanuel", LastName: "Thuring"}) // non-zero fields
	// db.Model(&user).Updates(map[string]interface{}{"first_name": "Olala", "last_name": "F42"})

	// Delete - delete user
	db.Delete(&user, 42)

	fmt.Printf("Serving on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
}
