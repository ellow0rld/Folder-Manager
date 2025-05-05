package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for files in the documents folder based on name or extension",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please specify a search term (name/extension) in Documents folder.")
			return
		}
		term := args[0]
		err := searchFiles(term)
		if err != nil {
			fmt.Println("Error searching files:", err)
		}
	},
}

func searchFiles(term string) error {
	dir := "C:/Users/<username>/Documents"  // <username> fill username with your system name
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Skipping:", path, "Error:", err)
			return nil
		}

		if filepath.Base(path) == term || filepath.Ext(path) == term {
			fmt.Println(path)
		}
		return nil
	})

	return err
}
