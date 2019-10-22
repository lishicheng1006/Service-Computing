package cmd

import (
	"errors"
	"github.com/Jiahonzheng/Go-Agenda/service"
	"github.com/spf13/cobra"
)

var (
	loginCmd = cobra.Command{
		Use:   "l",
		Short: "Log in",
		Long:  "Log in Agenda with username and password of a registered account",
	}
	loginCmdUsernameP = loginCmd.Flags().StringP("username", "u", "", "username of a registered account")
	loginCmdPasswordP = loginCmd.Flags().StringP("password", "p", "", "password of a registered account")
)

func login(cmd *cobra.Command, args []string) error {
	if *loginCmdUsernameP == "" {
		return errors.New("username is required")
	}
	if *loginCmdPasswordP == "" {
		return errors.New("password is required")
	}
	err := service.Login(*loginCmdUsernameP, *loginCmdPasswordP)
	checkError(err)
	return nil
}

var (
	logoutCmd = cobra.Command{
		Use:   "o",
		Short: "Log out",
		Long:  "Log out current account",
	}
)

func logout(cmd *cobra.Command, args []string) error {
	err := service.Logout()
	checkError(err)
	return nil
}

func init() {
	loginCmd.RunE = login
	logoutCmd.RunE = logout
	rootCmd.AddCommand(&loginCmd, &logoutCmd)
}
