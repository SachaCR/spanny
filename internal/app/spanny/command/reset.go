package command

import (
	"fmt"
	"slices"
	"sort"

	"github.com/SachaCR/spanny/internal/pkg/migration"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all migrations",
	Run: func(cmd *cobra.Command, args []string) {

		lastMigrationApplied, err := migration.GetLastMigrationApplied(getDatabasePath())

		if err != nil {
			fmt.Printf("Error getting last migration applied: %s\n", err)
			return
		}

		migrationList, err := migration.ListFiles(config.MigrationFilesPath)
		if err != nil {
			fmt.Printf("Error getting migration list: %s\n", err)
			return
		}

		if len(migrationList) == 0 {
			fmt.Println("No migration found")
			return
		}

		sort.Strings(migrationList)

		migrationIndex := slices.Index(migrationList, lastMigrationApplied)

		if migrationIndex == -1 {
			println("No migration to rollback")
			return
		}

		migrationNamesToRollback := migrationList[0 : migrationIndex+1]
		slices.Reverse(migrationNamesToRollback)

		for _, migrationName := range migrationNamesToRollback {
			print("Rollback migration ⏪: ", migrationName)

			err = migration.RunMigration(migration.ApplyMigrationParams{
				Direction:          migration.Down,
				DatabasePath:       getDatabasePath(),
				MigrationName:      migrationName,
				MigrationFilesPath: config.MigrationFilesPath,
			})

			if err != nil {
				println(" ❌")
				fmt.Printf("Error while migration rollback: %s\n", err)
				return
			}

			println(" ✅")
		}
	},
}
