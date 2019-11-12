package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// checkError is a tool function helps handle error.
func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "Application for meeting management.",
	Long:  "Agenda is an application for meeting management.",
}

// Execute runs the CLI application.
func Execute() {
	err := rootCmd.Execute()
	checkError(err)
}
