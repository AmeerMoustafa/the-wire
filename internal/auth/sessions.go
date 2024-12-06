package auth

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Username string
	Expiry   time.Time
}

var Sessions = map[string]Session{}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func GenerateSession(username string) http.Cookie {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	Sessions[sessionToken] = Session{
		Username: username,
		Expiry:   expiresAt,
	}

	sessionCookie := http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	}

	return sessionCookie
}
