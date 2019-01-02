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
	name := c.FormValue("name")
	var user User
	this.db.First(&user, "Name = ?", name)
	return c.String(http.StatusOK, fmt.Sprintln(user.Password))
}
