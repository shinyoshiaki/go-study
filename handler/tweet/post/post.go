package post

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"echo-pg/handler/user/login"
	"echo-pg/handler/user/utill"
	"echo-pg/model/tweet"
)

var (
	db    = tweet.Connect()
	count = -1
)

func Post(c echo.Context) (err error) {

	var json struct {
		Session string `json:"session"`
		Code    string `json:"code"`
		Text    string `json:"text"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	if login.IsLogin(c, json.Code, json.Session) == false {
		fmt.Println("not login")
		var result struct {
			State string `json:"state"`
		}
		result.State = "not login"
		return c.JSON(http.StatusBadRequest, result)
	}

	var tweets []tweet.Tweet
	if count == -1 {
		db.Find(&tweets).Count(count)
		if count == -1 {
			count = 0
		}
	}
	const layout = "2006/1/2 15:04:05"
	now := time.Now().Format(layout)

	db.Create(&tweet.Tweet{Number: count, Time: now, Code: json.Code, Text: json.Text})

	var result struct {
		Number int    `json:"number"`
		Time   string `json:"time"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Text   string `json:"text"`
	}

	result.Number = count

	result.Time = now
	result.Name = utill.Code2Name(json.Code)
	result.Code = json.Code
	result.Text = json.Text

	count++

	return c.JSON(http.StatusOK, result)
}
