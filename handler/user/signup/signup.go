package signup

import (
	"fmt"
	"net/http"
	"strconv"

	"echo-pg/handler/user/login"
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

	if err = c.Bind(&json); err != nil {
		return
	}

	if json.Name == "" || json.Password == "" {
		fmt.Println("null")
		return c.String(http.StatusBadRequest, fmt.Sprintln("null"))
	}

	var u user.User
	db.Find(&u, "Name = ?", json.Name)

	if u.Name == "" {
		fmt.Println("un exist")
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

		login.WriteCookie(c, code)

		return c.JSON(http.StatusOK, result)
	}
	fmt.Println("exist")
	return c.String(http.StatusBadRequest, "exist")
}
