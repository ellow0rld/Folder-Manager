package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move <source file path> <destination folder path>",
	Short: "Move a file to another directory",
	Long: `Move a file from one location to another.

Usage Example:
  CliProj move C:\Users\<username>\Downloads\file.txt C:\Users\<username>\Documents

If the destination is a folder, the file will be placed inside it.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]
		dest := args[1]

		err := moveFile(src, dest)
		if err != nil {
			fmt.Println("Error moving file:", err)
		} else {
			fmt.Println("File moved successfully!")
		}
	},
}

func moveFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("unable to open source file: %v", err)
	}
	defer sourceFile.Close()

	destInfo, err := os.Stat(dest)
	if err == nil && destInfo.IsDir() {
		dest = filepath.Join(dest, filepath.Base(src))
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("unable to create destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("unable to copy file: %v", err)
	}

	sourceFile.Close()
	destFile.Close()

	err = os.Remove(src)
	if err != nil {
		fmt.Println("Warning: unable to remove source file:", err)
	} else {
		fmt.Println("Source file deleted successfully.")
	}

	logAction("move", src, dest)
	return nil
}
