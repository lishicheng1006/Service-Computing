package cmd

import (
	"errors"
	"github.com/Jiahonzheng/Go-Agenda/service"
	"github.com/spf13/cobra"
)

// Register Command
var (
	registerCmd = cobra.Command{
		Use:   "r",
		Short: "Register a new user",
		Long:  "Register a new user with username, password, email and phone",
	}
	registerUsernameP = registerCmd.Flags().StringP("username", "u", "", "username of the user")
	registerPasswordP = registerCmd.Flags().StringP("password", "p", "", "password of the user")
	registerEmailP    = registerCmd.Flags().StringP("email", "e", "", "email of the user")
	registerPhoneP    = registerCmd.Flags().StringP("phone", "t", "", "phone of the user")
)

// register responds to Register Command.
func register(cmd *cobra.Command, args []string) error {
	if *registerUsernameP == "" {
		return errors.New("username is required")
	}
	if *registerPasswordP == "" {
		return errors.New("password is required")
	}
	if *registerEmailP == "" {
		return errors.New("email is required")
	}
	if *registerPhoneP == "" {
		return errors.New("phone is required")
	}
	err := service.Register(*registerUsernameP, *registerPasswordP, *registerEmailP, *registerPhoneP)
	checkError(err)
	return nil
}

// List All Users Command
var (
	listAllUsersCmd = cobra.Command{
		Use:   "lau",
		Short: "List all users",
		Long:  "List all users, logged in required",
	}
)

// listAllUsers responds to List All Users Command.
func listAllUsers(cmd *cobra.Command, args []string) error {
	err := service.ListAllUsers()
	checkError(err)
	return nil
}

// Delete Current User Command
var (
	deleteCurrentUserCmd = cobra.Command{
		Use:   "dc",
		Short: "Delete current user",
		Long:  "Delete current user, logged in required",
	}
)

// deleteCurrentUser responds to Delete Current User Command.
func deleteCurrentUser(cmd *cobra.Command, args []string) error {
	err := service.DeleteCurrentUser()
	checkError(err)
	return nil
}

func init() {
	registerCmd.RunE = register
	listAllUsersCmd.RunE = listAllUsers
	deleteCurrentUserCmd.RunE = deleteCurrentUser
	// Add commands.
	rootCmd.AddCommand(&registerCmd, &listAllUsersCmd, &deleteCurrentUserCmd)
}
