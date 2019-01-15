package main

import (
	addFollow "echo-pg/handler/follow/add"
	postTweet "echo-pg/handler/tweet/post"
	"echo-pg/handler/user/login"
	"echo-pg/handler/user/setting"
	"echo-pg/handler/user/signup"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
	e.Use(session.Middleware(store))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/users/signup", signup.SignUp)
	e.POST("/users/login", login.Login)
	e.POST("/users/update", setting.Update)
	e.POST("/tweet/post", postTweet.Post)
	e.POST("/follow/add", addFollow.Follow)
	// e.DELETE("/users/:name", Handler.DeleteUser)
	e.Start(":1323")
}

/*
curl -X POST localhost:1323/users -H 'Content-Type: application/json' -d '{"name":"update-work","password":"1234"}'
create table User(Name,Password,Key,Code);
create table Tweet(Number,Name,Code,Time,Text,Picture);
*/
