package post

import (
	"net/http"
	"strings"

	"echo-pg/model/follow"

	"github.com/labstack/echo"
)

var (
	db = follow.Connect()
)

func Info(c echo.Context) (err error) {
	var json struct {
		Code string `json:"code"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	var f follow.Follow
	db.First(&f, "Code = ?", json.Code)
	listFollow := strings.Split(f.Follow, "%")
	listFollower := strings.Split(f.Follower, "%")

	var result struct {
		Follow   []string `json:"follow"`
		Follower []string `json:"follower"`
	}
	result.Follow = listFollow
	result.Follower = listFollower

	return c.JSON(http.StatusOK, result)
}
