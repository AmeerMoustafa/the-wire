package auth

import (
	"time"
)

type Session struct {
	Username string
	Expiry   time.Time
}

var Sessions = map[string]Session{}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
