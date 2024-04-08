package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type MigrationType string

const (
	DDL MigrationType = "DDL"
	DML MigrationType = "DML"
)

var createDDLCmd = &cobra.Command{
	Use:   "create-ddl [migration-name]",
	Short: "Create migration files for DDL modification",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		createMigrationFiles(fileName, DDL)
	},
}

var createDMLCmd = &cobra.Command{
	Use:   "create-dml [migration-name]",
	Short: "Create migration files for DML modification",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		createMigrationFiles(fileName, DML)
	},
}

func createMigrationFiles(fileName string, migrationType MigrationType) {
	// Get current time
	currentTime := time.Now()

	migrationName := currentTime.Format("20060102150405123") + "_DML_" + fileName

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
