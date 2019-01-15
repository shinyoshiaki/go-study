package follow

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Follow struct {
	gorm.Model
	Code     string
	Follow   string //json配列をぶち込む
	Follower string
}

// Connect .
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(Follow{})
	return db
}
