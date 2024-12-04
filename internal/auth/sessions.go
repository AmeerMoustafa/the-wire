package auth

import (
	"time"
)

type Session struct {
	Username string
	Expiry   time.Time
}

var Sessions = map[string]Session{}

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}
