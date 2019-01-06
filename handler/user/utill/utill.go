package utill

import (
	"echo-pg/model/user"
)

var (
	db = user.Connect()
)

// UserNum .
func UserNum() int {
	count := 0
	var u []user.User
	db.Find(&u).Count(&count)
	return count
}
