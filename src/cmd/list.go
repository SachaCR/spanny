package cmd

import (
	"fmt"
	"spanny/src/migration"
	"spanny/src/ui"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List migration files",
	Run: func(cmd *cobra.Command, args []string) {
		fileList, err := migration.ListFiles(config.MigrationFilesPath)
		if err != nil {
			println(err.Error())
			return
		}

		columns := []string{
			"Migration files",
		}

		var rows [][]string
		for _, file := range fileList {
			rows = append(rows, []string{file})
		}

		if verbose {
			fmt.Println(config.MigrationFilesPath, ":")
		}

		ui.RenderTable(columns, rows)
	},
}
