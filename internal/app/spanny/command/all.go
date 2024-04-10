package command

import (
	"fmt"
	"slices"
	"sort"

	"github.com/SachaCR/spanny/internal/pkg/migration"
	"github.com/spf13/cobra"
)

var applyAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Apply all available migrations",
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

		if migrationIndex == len(migrationList)-1 {
			println("Already up to date. No migration to apply")
			return
		}

		migrationNamesToApply := migrationList[migrationIndex+1:]

		for _, migrationName := range migrationNamesToApply {
			print("Applying migration ⏩: ", migrationName)

			err = migration.RunMigration(migration.ApplyMigrationParams{
				Direction:          migration.Up,
				DatabasePath:       getDatabasePath(),
				MigrationName:      migrationName,
				MigrationFilesPath: config.MigrationFilesPath,
			})

			if err != nil {
				println(" ❌")
				fmt.Printf("Error applying migration: %s\n", err)
				return
			}

			println(" ✅")
		}
	},
}
