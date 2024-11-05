package common

import (
	"NoteAssistant/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {

}

var DB *gorm.DB

func InitDataBase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	initTables()
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitDataBase()
	}
	return DB
}

func initTables() {
	DB.AutoMigrate(model.User{})
}
