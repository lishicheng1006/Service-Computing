package entity

import (
	"log"
	"os"
)

// Session maintains Login Session.
type Session struct {
	Username string `json:"username"`
}

var sessionStorage = Storage{Path: "session.json"}

// GetSession return the current session.
func GetSession() *Session {
	session := new(Session)
	// If session.json exists, load it.
	if err := sessionStorage.Load(session); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Unable to load session data.")
	}
	return session
}

// GetCurrentUser returns the current user logged in.
func (s *Session) GetCurrentUser() (string, bool) {
	if s.Username == "" {
		return "", false
	}
	return s.Username, true
}

// Login sets the session.
func (s *Session) Login(username string) error {
	s.Username = username
	return sessionStorage.Save(s)
}

// Logout resets the session.
func (s *Session) Logout() error {
	s.Username = ""
	return sessionStorage.Save(s)
}
