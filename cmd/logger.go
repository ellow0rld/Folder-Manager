package cmd

import (
	"fmt"
	"os"
	"time"
)

// logAction records file movements for undo functionality
func logAction(action, oldPath, newPath string) error {
	logEntry := fmt.Sprintf("%s, %s, %s, %s\n", time.Now().Format(time.RFC3339), action, oldPath, newPath)

	f, err := os.OpenFile("history.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(logEntry); err != nil {
		return fmt.Errorf("error writing log: %w", err)
	}

	return nil
}
