package entity

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Users struct {
	Users []*User
}

var usersStorage = Storage{Path: "users.json"}

func GetUsers() *Users {
	users := new(Users)
	if err := usersStorage.Load(&users.Users); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Unable to load users data.")
	}
	return users
}

func (us *Users) Save() error {
	return usersStorage.Save(us.Users)
}

func (us *Users) Add(u *User) bool {
	if _, ok := us.FindByUsername(u.Username); ok {
		return false
	}
	us.Users = append(us.Users, u)
	return true
}

func (us *Users) FindByUsername(username string) (*User, bool) {
	for _, u := range us.Users {
		if u.Username == username {
			return u, true
		}
	}
	return nil, false
}

func (us *Users) DeleteByUsername(username string) error {
	users := us.Users
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i], users[len(users)-1] = users[len(users)-1], users[i]
			us.Users = users[:len(users)-1]
			return nil
		}
	}
	return fmt.Errorf("user %s not found", username)
}

func (us *Users) ListUsers() []*User {
	return us.Users
}
