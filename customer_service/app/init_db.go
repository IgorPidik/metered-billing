package main

import (
	"customer_service/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Service{})

	return db, nil
}
