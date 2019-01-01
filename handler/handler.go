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

type User struct {
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"password"`
}

type GetUser struct {
	Name string `json:"name" form:"name" query:"name"`
}

type Product struct {
	gorm.Model
	Name     string
	Password string
}

func Init(db *gorm.DB) Handler {
	return Handler{db: db}
}

func (this Handler) User(c echo.Context) (err error) {
	u := new(User)

	if err = c.Bind(u); err != nil {
		return
	}

	this.db.Create(&Product{Name: u.Name, Password: u.Password})
	return c.JSON(http.StatusOK, u)
}

func (this Handler) GetUser(c echo.Context) (err error) {
	g := new(GetUser)
	if c.Bind(g); err != nil {
		return
	}
	var product Product
	this.db.First(&product, "Name = ?", g.Name)
	return c.String(http.StatusOK, fmt.Sprint(product.Password))
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
