package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CliProj",
	Short: "File management CLI tool",
	Long:  `A command-line tool for organizing, moving, and searching files.`,
}

func init() {
	// Adding commands to rootCmd
	rootCmd.AddCommand(organizeCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(undoCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// Chatbot interpreter
