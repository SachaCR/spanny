package command

import (
	"fmt"

	"github.com/SachaCR/spanny/internal/pkg/migration"

	"github.com/spf13/cobra"
)

var initMigrationCmd = &cobra.Command{
	Use:   "init-migration",
	Short: "Will create the migration utility table in the database",
	Run: func(cmd *cobra.Command, args []string) {
		projectId := config.ProjectId
		instanceId := config.InstanceId
		databaseId := config.DatabaseId

		databasePath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, databaseId)

		err := migration.CreateMigrationTables(databasePath)

		if err != nil {
			fmt.Printf("Error creating migration tables: %s\n", err)
			return
		}

		err = migration.InsertInitialLockRow(databasePath)
		if err != nil {
			fmt.Printf("Error configuring migration lock: %s\n", err)
			return
		}

		fmt.Printf("Migration tables created\n")
	},
}
