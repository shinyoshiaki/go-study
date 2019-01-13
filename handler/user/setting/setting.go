package setting

import (
	"fmt"
	"net/http"

	"echo-pg/handler/user/login"
	"echo-pg/model/user"
	"echo-pg/utill/hash"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

var (
	db = user.Connect()
)

func Update(c echo.Context) (err error) {
	sess, _ := session.Get("session", c)
	fmt.Println("test", sess.Values["test"])

	var json struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Session  string `json:"session"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	fmt.Println(json)

	if login.IsLogin(c, json.Code, json.Session) == false {
		fmt.Println("not login")
		return c.String(http.StatusBadRequest, "not login")
	}

	if json.Name == "" || json.Password == "" {
		fmt.Println("null")
		return c.String(http.StatusBadRequest, "null")
	}

	var prev user.User
	db.Find(&prev, "Name = ?", json.Name)

	pass := hash.Sha1(json.Password)
	next := &user.User{Name: json.Name, Password: pass, Key: prev.Key, Code: prev.Code}
	db.Model(&prev).Update(next)

	var result struct {
		Name string `json:"name"`
		Code string `json:"id"`
	}
	result.Name = json.Name
	result.Code = json.Code

	return c.JSON(http.StatusOK, result)
}
