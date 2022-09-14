package utils

import (
	"billing_service/app/models"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.APIHit{})
	db.AutoMigrate(&models.Invoice{})
	// log.Println(e)

	return db, nil
}

func CreateTestData(db *gorm.DB) {
	log.Println("creating data")
	db.Create(&models.APIHit{UUID: uuid.New(), CustomerID: 1, ServiceID: 1, Timestamp: time.Now()})
	db.Create(&models.APIHit{UUID: uuid.New(), CustomerID: 1, ServiceID: 2, Timestamp: time.Now()})
	db.Create(&models.APIHit{UUID: uuid.New(), CustomerID: 2, ServiceID: 5, Timestamp: time.Now()})
}
