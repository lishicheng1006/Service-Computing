package entity

import (
	"fmt"
	"log"
	"os"
)

// User struct.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// Users maintains the list of User.
type Users struct {
	Users []*User
}

var usersStorage = Storage{Path: "users.json"}

// Return the current users.
func GetUsers() *Users {
	users := new(Users)
	if err := usersStorage.Load(&users.Users); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Unable to load users data.")
	}
	return users
}

// Save persists the current users data.
func (us *Users) Save() error {
	return usersStorage.Save(us.Users)
}

// Add a new user.
func (us *Users) Add(u *User) bool {
	if _, ok := us.FindByUsername(u.Username); ok {
		return false
	}
	us.Users = append(us.Users, u)
	return true
}

// FindByUsername checks whether an user exists.
func (us *Users) FindByUsername(username string) (*User, bool) {
	for _, u := range us.Users {
		if u.Username == username {
			return u, true
		}
	}
	return nil, false
}

// DeleteByUsername deletes an user.
func (us *Users) DeleteByUsername(username string) error {
	users := us.Users
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			// Delete the user from the list.
			users[i], users[len(users)-1] = users[len(users)-1], users[i]
			us.Users = users[:len(users)-1]
			return nil
		}
	}
	return fmt.Errorf("user %s not found", username)
}

// ListAllUsers list all users.
func (us *Users) ListAllUsers() []*User {
	return us.Users
}
