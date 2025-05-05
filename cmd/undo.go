package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// undoCmd defines the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last action",
	Run: func(cmd *cobra.Command, args []string) {
		err := undoLastAction()
		if err != nil {
			fmt.Println("Error undoing last action:", err)
		} else {
			fmt.Println("Undo successful!")
		}
	},
}

// undoLastAction reverts the last file move or organize operation
func undoLastAction() error {
	historyFile := "history.log"

	// Read history file
	lines, err := readLines(historyFile)
	if err != nil {
		return fmt.Errorf("could not read history file: %v", err)
	}
	if len(lines) == 0 {
		return fmt.Errorf("no history found")
	}

	// Get the last action line
	lastAction := lines[len(lines)-1]
	parts := strings.Split(lastAction, ", ")

	// Ensure the line format is correct
	if len(parts) < 4 {
		return fmt.Errorf("invalid history format: %s", lastAction)
	}

	// Extract details from history
	action := strings.TrimSpace(parts[1])
	sourcePath := strings.TrimSpace(parts[2])
	destinationPath := strings.TrimSpace(parts[3])

	// Remove last log entry
	err = writeLines(lines[:len(lines)-1], historyFile)
	if err != nil {
		return fmt.Errorf("error updating history file: %v", err)
	}

	// Handle undo based on action type
	switch action {
	case "move", "organize":
		err = movesFile(destinationPath, sourcePath)
		if err != nil {
			fmt.Println("Error moving file back:", err)
		} else {
			fmt.Printf("Restored '%s' back to '%s' (Undo %s)\n", destinationPath, sourcePath, action)
		}
	default:
		return fmt.Errorf("unsupported action type in history: %s", action)
	}

	return nil
}

// moveFile moves a file from src to dst
func movesFile(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return fmt.Errorf("failed to move file: %v", err)
	}
	return nil
}

// readLines reads all lines from a file
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes lines to a file
func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
