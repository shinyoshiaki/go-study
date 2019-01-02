package handler

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Handler struct {
	db *gorm.DB
}

type UserJson struct {
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
}

type GetUser struct {
	Name string `json:"name" form:"name" query:"name"`
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

func Init(db *gorm.DB) Handler {
	return Handler{db: db}
}

func (this Handler) CreateUser(c echo.Context) (err error) {
	u := new(UserJson)

	if err = c.Bind(u); err != nil {
		return
	}

	this.db.Create(&User{Name: u.Name, Password: u.Password})
	return c.JSON(http.StatusOK, u)
}

func (this Handler) GetUser(c echo.Context) (err error) {
	name := c.Param("name")
	var user User
	this.db.First(&user, "Name = ?", name)

	fmt.Println(user.Name + ":" + user.Password)

	return c.String(http.StatusOK, fmt.Sprintln(user.Password))
}

func (this Handler) UpdateUser(c echo.Context) (err error) {
	name := c.Param("name")
	var prev User
	this.db.First(&prev, "Name = ?", name)

	fmt.Println(prev.Name + ":" + prev.Password)

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
