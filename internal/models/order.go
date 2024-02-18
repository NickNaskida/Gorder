package models

import (
	"errors"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderId     string `gorm:"unique;not null;index"`
	Name        string `gorm:"not null"`
	Description string
	Price       int32 `gorm:"not null"`
}

func (order *Order) Validate() error {
	if order.Name == "" || order.Price == 0 {
		return errors.New("name and price are required fields")
	}

	if order.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}
