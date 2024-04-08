package cmd

import (
	"fmt"
	"spanny/src/dbops"
	"spanny/src/migration"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Will create the instance, database and utility migration tables",
	Run: func(cmd *cobra.Command, args []string) {
		projectId := config.ProjectId
		instanceId := config.InstanceId
		databaseId := config.DatabaseId

		databasePath := getDatabasePath()

		_, err := dbops.CreateInstance(projectId, instanceId)
		if err != nil {
			fmt.Printf("Error creating instance: %s\n", err)
			return
		}
		fmt.Printf("Instance created: %s\n", instanceId)

		_, err = dbops.CreateDatabase(instanceId, projectId, databaseId)
		if err != nil {
			fmt.Printf("Error creating database: %s\n", err)
			return
		}
		fmt.Printf("Database created: %s\n", databaseId)

		err = migration.CreateMigrationTables(databasePath)
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
