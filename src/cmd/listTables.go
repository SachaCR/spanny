package cmd

import (
	"spanny/src/dbops"
	"spanny/src/ui"

	"github.com/spf13/cobra"
)

var listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "List tables files",
	Run: func(cmd *cobra.Command, args []string) {
		databasePath := getDatabasePath()
		columns, rows, err := dbops.ExecuteReadQuery(databasePath, `SELECT table_name from information_schema.tables WHERE table_type = "BASE TABLE"`)

		if err != nil {
			println(err.Error())
			return
		}

		ui.RenderTable(columns, rows)
	},
}
