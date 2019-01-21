package post

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"echo-pg/handler/user/login"
	"echo-pg/handler/user/utill"
	"echo-pg/model/follow"
	"echo-pg/utill/array"
)

var (
	db = follow.Connect()
)

func Follow(c echo.Context) (err error) {
	var json struct {
		Session string `json:"session"`
		Code    string `json:"code"`
		User    string `json:"user"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	if login.IsLogin(c, json.Code, json.Session) == false {
		return c.String(http.StatusBadRequest, "not login")
	}

	if utill.Code2Name(json.Code) == json.User == true {
		return c.String(http.StatusBadRequest, "its me")
	}

	var f follow.Follow
	db.First(&f, "Code = ?", json.Code)
	list := strings.Split(f.Follow, "%")
	if array.Includes(list, json.User) == true {
		return c.String(http.StatusBadRequest, "exist")
	}
	var follows string
	if f.Follow == "" {
		follows = json.User
	} else {
		follows = f.Follow + "%" + json.User
	}
	db.Model(&f).Update(&follow.Follow{Follow: follows})

	var u follow.Follow
	db.First(&u, "Code = ?", json.User)
	var followers string
	if u.Follow == "" {
		followers = json.User
	} else {
		followers = u.Follow + "%" + json.User
	}
	db.Model(&u).Update(&follow.Follow{Follower: followers})

	var result struct {
		User string `json:"user"`
	}

	result.User = json.User

	return c.JSON(http.StatusOK, result)
}
