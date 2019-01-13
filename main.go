package main

import (
	"echo-pg/handler/user/login"
	"echo-pg/handler/user/setting"
	"echo-pg/handler/user/signup"

	"github.com/gorilla/sessions"
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// maxAge := 24 * 60 * 60
	maxAge := 60 * 10

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/users/signup", signup.SignUp)
	e.POST("/users/login", login.Login)
	e.POST("/users/update", setting.Update)
	// e.DELETE("/users/:name", Handler.DeleteUser)
	e.Start(":1323")
}

/*
curl -X POST localhost:1323/users -H 'Content-Type: application/json' -d '{"name":"update-work","password":"1234"}'
curl localhost:1323/users/update-work
curl -X PUT http://localhost:1323/users/update-work -H 'Content-Type: application/json' -d '{"password":"password"}'
curl -X DELETE localhost:1323/users/update-work
create table User(Name,Password,Key,Code);
*/
