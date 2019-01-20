package get

import (
	"net/http"

	"echo-pg/handler/user/utill"
	"echo-pg/model/tweet"

	"github.com/labstack/echo"
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

	type Tweet struct {
		Number int    `json:"number"`
		Time   string `json:"time"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Text   string `json:"text"`
	}
	results := make([]Tweet, 0)

	for _, v := range tweets {
		tweet := Tweet{}
		tweet.Number = v.Number
		tweet.Time = v.Time
		tweet.Name = utill.Code2Name(v.Code)
		tweet.Code = v.Code
		tweet.Text = v.Text
		results = append(results, tweet)
	}

	var result struct {
		Tweets []Tweet `json:"tweets"`
	}
	result.Tweets = results
	return c.JSON(http.StatusOK, result)
}

func Search(c echo.Context) (err error) {
	var json struct {
		Word string `json:"word"`
	}
	if err = c.Bind(&json); err != nil {
		return
	}

	tweets := []tweet.Tweet{}
	db.Where("Text LIKE ?", "%"+json.Word+"%").Find(&tweets)

	type Tweet struct {
		Number int    `json:"number"`
		Time   string `json:"time"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Text   string `json:"text"`
	}
	results := make([]Tweet, 0)

	for _, v := range tweets {
		tweet := Tweet{}
		tweet.Number = v.Number
		tweet.Time = v.Time
		tweet.Name = utill.Code2Name(v.Code)
		tweet.Code = v.Code
		tweet.Text = v.Text
		results = append(results, tweet)
	}

	var result struct {
		Tweets []Tweet `json:"tweets"`
		Word   string  `json:"word"`
	}
	result.Tweets = results
	result.Word = json.Word
	return c.JSON(http.StatusOK, result)
}
