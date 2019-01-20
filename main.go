package main

import (
	addFollow "echo-pg/handler/follow/add"
	getFollow "echo-pg/handler/follow/get"
	getTweet "echo-pg/handler/tweet/get"
	postTweet "echo-pg/handler/tweet/post"
	"echo-pg/handler/user/login"
	"echo-pg/handler/user/setting"
	"echo-pg/handler/user/signup"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/users/signup", signup.SignUp)
	e.POST("/users/login", login.Login)
	e.POST("/users/update", setting.Update)
	e.POST("/tweet/post", postTweet.Post)
	e.POST("/tweet/get/mine", getTweet.Mine)
	e.POST("/tweet/get/search", getTweet.Search)
	e.POST("/tweet/get/user", getTweet.User)
	e.POST("/follow/add", addFollow.Follow)
	e.POST("/follow/get", getFollow.Info)
	// e.DELETE("/users/:name", Handler.DeleteUser)
	e.Start(":1323")
}

/*
curl -X POST localhost:1323/users -H 'Content-Type: application/json' -d '{"name":"update-work","password":"1234"}'
create table User(Name,Password,Key,Code);
create table Tweet(Number,Name,Code,Time,Text,Picture);
*/
