package utill

import (
	"echo-pg/model/follow"
)

var (
	db = follow.Connect()
)

func Init(code string) {
	var f follow.Follow
	db.First(&f, "Code = ?", code)
	if f.Code == "" {
		db.Create(&follow.Follow{Code: code})
	}
}
