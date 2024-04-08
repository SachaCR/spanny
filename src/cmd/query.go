package cmd

import (
	"spanny/src/dbops"
	"spanny/src/ui"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run a read only query against the database",
	Run: func(cmd *cobra.Command, args []string) {
		query := ui.OpenTextArea()

		if query == "" {
			return
		}

		databasePath := getDatabasePath()

		columns, rows, err := dbops.ExecuteReadQuery(databasePath, query)
		if err != nil {
			println(err.Error())
			return
		}

		if len(rows) == 0 {
			println("No results found.")
			return
		}

		ui.RenderTable(columns, rows)
	},
}
