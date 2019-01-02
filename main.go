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
	e.POST("/users", Handler.CreateUser)
	e.GET("/users/:name", Handler.GetUser)
	e.PUT("/users/:name", Handler.UpdateUser)
	e.DELETE("/users/:name", Handler.DeleteUser)
	e.Start(":1323")
}

/*
curl -X POST localhost:1323/users -H 'Content-Type: application/json' -d '{"name":"update-work","password":"1234"}'
curl localhost:1323/users/update-work
curl -X PUT http://localhost:1323/users/update-work -H 'Content-Type: application/json' -d '{"password":"password"}'
curl -X DELETE localhost:1323/users/update-work
*/
