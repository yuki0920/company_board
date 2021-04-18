package models

import "gorm.io/gorm"

// メソッドにCountとTakeをもつstruct
type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, Limit int, offset int) interface{}
}
