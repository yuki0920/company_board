package database

import (
	"github.com/yuki0920/company_board/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// databaseを代入し、他パッケージから参照するための変数
var DB *gorm.DB

func Connect() {
	// 環境変数化したい
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "company_board"
	DSN := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	database, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = database

	// テーブルがうまく生成されないときは、引数を減らしてみるとうまくいくことがある
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Order{})
	database.AutoMigrate(&models.OrderItem{})
}
