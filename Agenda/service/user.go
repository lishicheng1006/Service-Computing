package service

import (
	"errors"
	"fmt"
	"github.com/Jiahonzheng/Go-Agenda/entity"
	"github.com/olekukonko/tablewriter"
	"os"
)

var ErrConflictUsername = errors.New("conflict username")

// Register a user via username, password, email and phone.
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

// ListAllUsers list all users.
func ListUsers() error {
	if _, err := checkLogin(); err != nil {
		return err
	}
	// Use TableWriter to display all users.
	table := tablewriter.NewWriter(os.Stdout)
	// Set the column names of the table.
	table.SetHeader([]string{"Username", "Email", "Phone"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, u := range users.ListAllUsers() {
		table.Append([]string{u.Username, u.Email, u.Phone})
	}
	// Print the table.
	table.Render()
	return nil
}

// DeleteCurrentUser deletes current user logged in.
func DeleteCurrentUser() error {
	username, err := checkLogin()
	if err != nil {
		return err
	}
	err = users.DeleteByUsername(username)
	if err != nil {
		return err
	}
	// Persist data.
	err = users.Save()
	if err != nil {
		return err
	}
	// Logout the current user.
	if err := session.Logout(); err != nil {
		return err
	}
	fmt.Printf("User %s has been deleted successfully.", username)
	return nil
}
