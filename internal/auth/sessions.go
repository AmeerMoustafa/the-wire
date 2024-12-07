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

func AuthenticateUser(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	sessionToken := cookie.Value

	userSession, exists := Sessions[sessionToken]
	if !exists {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}

	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}

	return true
}
