package session

import "time"

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
		time: time.Now(),
	}
	sessions[id] = s
}

func Get(id string) string {
	s := sessions[id]
	if s == nil {
		return ""
	}
	return s.key
}
