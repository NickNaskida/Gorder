package database

import (
	"gorm.io/driver/sqlite"
	"log"

	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
