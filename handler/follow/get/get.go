package post

import (
	"strings"

	"echo-pg/model/follow"
)

var (
	db = follow.Connect()
)

func Follow(code string) []string {
	var f follow.Follow
	db.First(&f, "Code = ?", code)
	list := strings.Split(f.Follow, "%")
	return list
}
