package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	URL string `gorm:"uniqueIndex"`
	Title       string
	Price       string
	Availability string
}