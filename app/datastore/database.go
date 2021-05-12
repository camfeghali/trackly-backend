package datastore

import (
	"fmt"
	"trackly-backend/app/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDBInstance(db_name, db_user, db_pw string) *DB {
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", db_name, db_user, db_pw)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.CheckError(err)
	return &DB{db}
}

func (db *DB) RunMigrations() error {
	err := db.AutoMigrate(&User{}, &Client{}, &Project{}, &Task{}, &TimeEntry{})
	return err
}