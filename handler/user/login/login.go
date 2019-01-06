package login

import (
	"net/http"

	"echo-pg/model/user"
	"echo-pg/utill/hash"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

var (
	db = user.Connect()
)

// Login login
func Login(c echo.Context) (err error) {

	var json struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err = c.Bind(json); err != nil {
		return
	}

	if json.Name == "" || json.Password == "" {
		return c.String(http.StatusBadRequest, "null")
	}

	var u user.User
	db.Find(&u, "Name = ?", json.Name)

	if u.Name == "" {
		return c.String(http.StatusBadRequest, "unexist")
	}

	if hash.Sha1(json.Password) == u.Password {
		session := session.Default(c)
		session.Set("loginCompleted", "completed")
		session.Save()

		var result struct {
			Name string `json:"name"`
			Code string `json:"id"`
		}
		result.Name = u.Name
		result.Code = u.Code
		return c.JSON(http.StatusOK, result)
	}
	return c.String(http.StatusBadRequest, "wrong pass")
}

// IsLogin .
func IsLogin(c echo.Context) bool {
	session := session.Default(c)

	login := session.Get("loginCompleted")
	if login != nil && login == "completed" {
		return true
	}
	return false
}
