package signup

import (
	"net/http"
	"strconv"

	"echo-pg/model/user"
	"echo-pg/utill/hash"

	"github.com/labstack/echo"
)

var (
	db = user.Connect()
)

// SignUp signup
func SignUp(c echo.Context) (err error) {

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
		count := 0
		var users []user.User
		db.Find(&users).Count(count)

		code := hash.Sha1(strconv.Itoa(count))
		pass := hash.Sha1(json.Password)

		db.Create(&user.User{Name: json.Name, Password: pass, Key: count, Code: code})
		var result struct {
			Name string `json:"name"`
			Code string `json:"id"`
		}
		result.Name = json.Name
		result.Code = code
		return c.JSON(http.StatusOK, result)
	}
	return c.String(http.StatusBadRequest, "exist")
}
