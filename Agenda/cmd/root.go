package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "CLI App for user and meeting management",
	Long:  "Agenda is an application for user and meeting management.",
}

func Execute() {
	err := rootCmd.Execute()
	checkError(err)
}
