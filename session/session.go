package session

import (
	"fmt"
	"time"
)

type (
	session struct {
		key  string
		time time.Time
	}
)

var sessions = map[string]*session{}

func Set(id string, key string) {
	s := &session{
		key:  key,
		time: time.Now().Add(time.Hour * 1),
	}
	sessions[id] = s
}

func Get(id string) string {
	s := sessions[id]
	if s == nil {
		fmt.Println("session null")
		return ""
	}
	if time.Now().Unix() > s.time.Unix() {
		fmt.Println("session timeout")
		return ""
	}
	return s.key
}
