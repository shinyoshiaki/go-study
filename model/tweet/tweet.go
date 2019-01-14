package tweet

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Tweet struct {
	gorm.Model
	Number  int
	Time    string
	Code    string
	Text    string
	Picture string
}

// Connect .
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(Tweet{})
	return db
}
