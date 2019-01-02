package main

import (
	"./db"
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	db := db.GormConnect()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Handler := handler.Init(db)
	e.POST("/save", Handler.CreateUser)
	e.POST("/get", Handler.GetUser)

	e.Start(":1323")
}
