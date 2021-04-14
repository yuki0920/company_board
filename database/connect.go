package database

import (
	"github.com/yuki0920/company_board/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// databaseを代入し、他パッケージから参照するための変数
var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}
