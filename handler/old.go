package old

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (this Handler) CreateUser(c echo.Context) (err error) {
	u := new(UserJson)

	if err = c.Bind(u); err != nil {
		return
	}

	var user User
	this.db.Find(&user, "Name = ?", u.Name)

	count := this.userNum()

	encoder := sha1.New()
	encoder.Write([]byte(strconv.Itoa(count)))
	hash := encoder.Sum(nil)
	hexHash := hex.EncodeToString(hash)

	if user.Name == "" {
		this.db.Create(&User{Name: u.Name, Password: u.Password, Key: count, Code: hexHash})
		var result struct {
			Name string `json:"name"`
			Code string `json:"id"`
		}
		result.Name = u.Name
		result.Code = hexHash
		return c.JSON(http.StatusOK, result)
	}
	return c.String(http.StatusBadRequest, "fail")
}

func (this Handler) GetUser(c echo.Context) (err error) {
	name := c.Param("name")
	var user User
	this.db.First(&user, "Name = ?", name)

	fmt.Println(user.Name + ":" + user.Password)

	return c.JSON(http.StatusOK, user)
}

func (this Handler) UpdateUser(c echo.Context) (err error) {
	name := c.Param("name")
	var prev User
	this.db.First(&prev, "Name = ?", name)

	if prev.Name == "" {
		return c.String(http.StatusBadRequest, "fail")
	}

	next := new(UserJson)
	if err = c.Bind(next); err != nil {
		return
	}
	next.Name = name

	fmt.Println(next.Password)
	this.db.Model(&prev).Update(&next)

	return c.String(http.StatusOK, fmt.Sprintln(next.Password))
}

func (this Handler) DeleteUser(c echo.Context) (err error) {
	name := c.Param("name")
	var u User
	this.db.First(&u, "Name = ?", name)
	this.db.Delete(&u)
	return c.JSON(http.StatusOK, u)
}

func (this Handler) userNum() int {
	count := 0
	var u []User
	this.db.Find(&u).Count(&count)
	return count
}

func (this Handler) GetUserNum(c echo.Context) (err error) {
	count := this.userNum()
	return c.String(http.StatusOK, fmt.Sprintln(count))
}
