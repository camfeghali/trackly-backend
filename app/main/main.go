package main

import "trackly-backend-gorm/app/database"

type User struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {
	db := database.NewDBInstance("trackly_user", "insabgho123", "trackly")
	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	db.Create(&User{FirstName: "D42", LastName: "Doodoo"})

	// Read
	var user User
	// db.First(&user)                            // find user with integer primary key
	db.First(&user, "last_name = ?", "Doodoo") // find user with code D42

	// Update - update user's price to 200
	db.Model(&user).Update("last_name", "Kathy")
	// Update - update multiple fields
	db.Model(&user).Updates(User{FirstName: "Emmanuel", LastName: "Thuring"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"first_name": "Olala", "last_name": "F42"})

	// Delete - delete user
	db.Delete(&user, 42)
}
