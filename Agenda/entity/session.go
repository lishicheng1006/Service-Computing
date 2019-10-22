package entity

import (
	"log"
	"os"
)

type Session struct {
	Username string `json:"username"`
}

var sessionStorage = Storage{Path: "session.json"}

func GetSession() *Session {
	session := new(Session)
	if err := sessionStorage.Load(session); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Unable to load session data.")
	}
	return session
}

func (s *Session) GetCurrentUser() (string, bool) {
	if s.Username == "" {
		return "", false
	}
	return s.Username, true
}

func (s *Session) Login(username string) error {
	s.Username = username
	return sessionStorage.Save(s)
}

func (s *Session) Logout() error {
	s.Username = ""
	return sessionStorage.Save(s)
}
