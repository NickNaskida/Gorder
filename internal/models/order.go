package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	OrderId     string `gorm:"unique;not null"`
	Name        string
	Description string
	Price       int32
}
