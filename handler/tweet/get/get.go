package get

import (
	"net/http"

	"github.com/labstack/echo"

	"echo-pg/model/tweet"
)

var (
	db = tweet.Connect()
)

func Mine(c echo.Context) (err error) {
	var json struct {
		Session string `json:"session"`
		Code    string `json:"code"`
	}
	if err = c.Bind(&json); err != nil {
		return
	}
	tweets := []tweet.Tweet{}
	db.Find(&tweets, "Code = ?", json.Code)
	var result struct {
		Tweets []tweet.Tweet
	}
	result.Tweets = tweets
	return c.JSON(http.StatusOK, result)
}
