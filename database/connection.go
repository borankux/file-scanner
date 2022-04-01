package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	opended, _ := gorm.Open(sqlite.Open("./files.db"))
	db = opended
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&File{})
}
