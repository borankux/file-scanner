package database

import (
	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	opended, _ := gorm.Open(sqlite.Open("./files.db"))
	db = opended
	if db.Error != nil {
		color.Red("failed:%s", db.Error.Error())
		return
	}
	db.AutoMigrate(&File{})
}

func GetDB() *gorm.DB {
	return db
}

func SaveFile(name string, size int64) {
	file := File{
		Name: name,
		Size: size,
	}
	db.Create(&file)
}
