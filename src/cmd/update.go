package cmd

import (
	"spanny/src/dbops"
	"spanny/src/ui"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Run insert or update against the database",
	Run: func(cmd *cobra.Command, args []string) {
		updateQuery := ui.OpenTextArea()

		if updateQuery == "" {
			return
		}

		databasePath := getDatabasePath()

		rowCount, err := dbops.ExecuteUpdateQuery(databasePath, updateQuery)

		if err != nil {
			println(err.Error())
			return
		}

		println("Rows affected: ", rowCount)
	},
}
