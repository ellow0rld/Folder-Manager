package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// organizeCmd defines the organize command
var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize your downloads folder",
	Run: func(cmd *cobra.Command, args []string) {
		downloadsDir := "C:/Users/madhu/Downloads" // Source directory
		destDir := "C:/Users/madhu/Documents"      // Destination directory where categorized folders exist

		fmt.Println("Organizing files...")
		err := organizeFiles(downloadsDir, destDir)
		if err != nil {
			fmt.Println("Error organizing files:", err)
		} else {
			fmt.Println("Files organized successfully!")
		}
	},
}

// getFolders retrieves existing folders from the destination directory
func getFolders(directory string) ([]string, error) {
	var folders []string
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}
	return folders, nil
}

// organizeFiles processes files and moves them based on topic classification
func organizeFiles(sourceDir, destDir string) error {
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	folders, err := getFolders(destDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(sourceDir, file.Name())

		// Extract text content
		text, err := extractText(filePath)
		var topic string

		if err != nil {
			// Prompt user for manual selection
			topic = userChooseTopic(folders, file.Name())
		} else {
			// Classify file based on content
			topic, err = classifyWithPython(text)
			if err != nil || topic == "" {
				topic = userChooseTopic(folders, file.Name()) // Ask user in case of ambiguity
			}
		}

		// Ensure target directory exists
		topicDir := filepath.Join(destDir, topic)
		if err := os.MkdirAll(topicDir, os.ModePerm); err != nil {
			fmt.Printf("Error creating folder %s: %v\n", topicDir, err)
			continue
		}

		// Move file to categorized folder
		newPath := filepath.Join(topicDir, file.Name())
		if err := os.Rename(filePath, newPath); err != nil {
			fmt.Printf("Error moving file %s: %v\n", file.Name(), err)
		} else {
			logAction("organize", filePath, newPath)
			fmt.Printf("Moved '%s' → [%s]\n", file.Name(), topic)
		}
	}
	return nil
}

// extractText extracts text from a file based on its type
func extractText(filePath string) (string, error) {
	fileExt := strings.ToLower(filepath.Ext(filePath))

	switch fileExt {
	case ".txt":
		// Read text files directly
		content, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("error reading text file: %w", err)
		}
		return string(content), nil

	case ".pdf":
		// Extract text from PDFs using Python script
		cmd := exec.Command("python", "extract_text.py", filePath)
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("error extracting text from PDF: %w", err)
		}
		return string(output), nil

	default:
		return "", fmt.Errorf("unsupported file type: %s", fileExt)
	}
}

// classifyWithPython runs a Python script to classify the file content
func classifyWithPython(text string) (string, error) {
	cmd := exec.Command("python", "classify.py")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close()

	// Send file content to Python
	go func() {
		_, _ = stdin.Write([]byte(text))
		stdin.Close()
	}()

	// Get output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse JSON response
	var result map[string]string
	if err := json.Unmarshal(output, &result); err != nil {
		return "", err
	}

	return result["topic"], nil
}

// userChooseTopic prompts the user to select a category manually, skip the file, or exit.
func userChooseTopic(folders []string, fileName string) string {
	fmt.Printf("\n📂 Ambiguous classification for '%s'. Choose a folder or take an action:\n", fileName)
	for i, folder := range folders {
		fmt.Printf("[%d] %s\n", i+1, folder)
	}
	fmt.Println("[S] Skip this file (Leave it in place)")
	fmt.Println("[Q] Quit (Stop organizing files)")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number, 'S' to skip, or 'Q' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		// If user chooses to skip, return an empty string (no move action)
		if input == "S" {
			return "" // Indicates skipping
		}

		// If user chooses to quit, exit the program
		if input == "Q" {
			fmt.Println("🛑 Organizing process stopped by user.")
			os.Exit(0)
		}

		// Validate numeric input for folder selection
		index, err := strconv.Atoi(input)
		if err == nil && index > 0 && index <= len(folders) {
			return folders[index-1]
		}

		fmt.Println("❌ Invalid choice. Please enter a valid number, 'S' to skip, or 'Q' to quit.")
	}
}
