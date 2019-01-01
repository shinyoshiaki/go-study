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

	e.GET("/hello", handler.MainPage())

	Post := handler.Init(db)
	e.POST("/save", Post.User)
	e.POST("/get", Post.GetUser)

	e.Start(":1323")
}
