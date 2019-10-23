package service

import (
	"errors"
	"fmt"
)

var ErrNeedLogIn = errors.New("need login")
var ErrHaveLoggedIn = errors.New("have logged in")
var ErrWrongUsernameOrPassword = errors.New("wrong username or password")

// checkLogin checks whether Agenda is logged in.
func checkLogin() (string, error) {
	username, ok := session.GetCurrentUser()
	if !ok {
		return "", ErrNeedLogIn
	}
	return username, nil
}

// Login via username and password.
func Login(username, password string) error {
	if _, err := checkLogin(); err == nil {
		return ErrHaveLoggedIn
	}
	if u, ok := users.FindByUsername(username); !ok || u.Password != password {
		return ErrWrongUsernameOrPassword
	}
	if err := session.Login(username); err != nil {
		return err
	}
	fmt.Printf("User %s has successfully logged in.", username)
	return nil
}

// Logout the current user.
func Logout() error {
	var username string
	username, err := checkLogin()
	if err != nil {
		return err
	}
	if err := session.Logout(); err != nil {
		return err
	}
	fmt.Printf("User %s has successfully logged out.", username)
	return nil
}
