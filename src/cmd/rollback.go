package cmd

import (
	"fmt"
	"sort"
	"spanny/src/migration"

	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback last migration",
	Run: func(cmd *cobra.Command, args []string) {

		lastMigrationApplied, err := migration.GetLastMigrationApplied(getDatabasePath())

		if err != nil {
			fmt.Printf("Error getting last migration applied: %s\n", err)
			return
		}
		if lastMigrationApplied == "NONE" {
			fmt.Println("No migration to rollback")
			return
		}

		migrationList, err := migration.ListFiles(config.MigrationFilesPath)

		if err != nil {
			fmt.Printf("Error getting migration list: %s\n", err)
			return
		}

		sort.Strings(migrationList)

		migrationNameToRollback := lastMigrationApplied

		print("Rollback migration ⏪: ", migrationNameToRollback)

		err = migration.RunMigration(migration.ApplyMigrationParams{
			Direction:          migration.Down,
			DatabasePath:       getDatabasePath(),
			MigrationName:      migrationNameToRollback,
			MigrationFilesPath: config.MigrationFilesPath,
		})

		if err != nil {
			println(" ❌")
			fmt.Printf("Error while migration rollback: %s\n", err)
			return
		}

		println(" ✅")
	},
}
