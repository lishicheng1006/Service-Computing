package service

import (
	"errors"
	"fmt"
	"github.com/Jiahonzheng/Go-Agenda/entity"
	"github.com/olekukonko/tablewriter"
	"os"
)

var ErrConflictUsername = errors.New("conflict username")

func Register(username, password, email, phone string) error {
	if ok := users.Add(&entity.User{Username: username, Password: password, Email: email, Phone: phone}); !ok {
		return ErrConflictUsername
	}
	if err := users.Save(); err != nil {
		return err
	}
	fmt.Printf("User %s has been registered successfully.", username)
	return nil
}

func ListUsers() error {
	if _, err := checkLogin(); err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Email", "Phone"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, u := range users.ListUsers() {
		table.Append([]string{u.Username, u.Email, u.Phone})
	}
	table.Render()
	return nil
}

func DeleteCurrentUser() error {
	username, err := checkLogin()
	if err != nil {
		return err
	}
	err = users.DeleteByUsername(username)
	if err != nil {
		return err
	}
	err = users.Save()
	if err != nil {
		return err
	}
	if err := session.Logout(); err != nil {
		return err
	}
	fmt.Printf("User %s has been deleted successfully.", username)
	return nil
}
