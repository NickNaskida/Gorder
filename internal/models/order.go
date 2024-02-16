package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	OrderId     string `gorm:"unique;not null;index"`
	Name        string `gorm:"not null"`
	Description string
	Price       int32 `gorm:"not null"`
}
