package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [migration-name]",
	Short: "Create migration files",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		createMigrationFiles(fileName)
	},
}

func createMigrationFiles(fileName string) {
	// Get current time
	currentTime := time.Now()

	migrationName := currentTime.Format("20060102150405123") + "_" + fileName

	// Create a folder with the timestamp as its name
	folderName := config.MigrationFilesPath + "/" + migrationName

	err := os.Mkdir(folderName, 0755)

	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	fmt.Println("Folder", folderName, "created successfully.")

	// Create files inside the folder
	upFilePath := filepath.Join(folderName, "up.sql")
	downFilePath := filepath.Join(folderName, "down.sql")

	// Create "up" file
	upFile, err := os.Create(upFilePath)
	if err != nil {
		fmt.Println("Error creating 'up' file:", err)
		return
	}
	defer upFile.Close()

	// Create "down" file
	downFile, err := os.Create(downFilePath)
	if err != nil {
		fmt.Println("Error creating 'down' file:", err)
		return
	}
	defer downFile.Close()

	fmt.Println("Files 'up.sql' and 'down.sql' created inside the folder.")
}
