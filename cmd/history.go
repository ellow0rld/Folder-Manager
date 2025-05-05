package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show history of actions performed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing history...")
		err := showHistory()
		if err != nil {
			fmt.Println("Error reading history:", err)
		}
	},
}

func showHistory() error {
	file, err := os.OpenFile("history.log", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open history log: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("History:")
	found := false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		found = true
	}

	if !found {
		fmt.Println("No history found.")
	}
	return scanner.Err()
}
