package database

import (
	"monitor-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("monitor.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}


	err = db.AutoMigrate(&models.Product{})
	return db, err
}