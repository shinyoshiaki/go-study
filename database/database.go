package database

import (
	"echo-pg/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func GormConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&handler.User{})
	return db
}
