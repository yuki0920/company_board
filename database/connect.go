package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(sqlite.Open("go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
}
