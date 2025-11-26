package command

import (
	"github.com/SachaCR/spanny/internal/pkg/dbops"
	"github.com/SachaCR/spanny/internal/pkg/ui"
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

		rowCount, err := dbops.ExecuteUpdateQuery(cmd.Context(), databasePath, updateQuery)

		if err != nil {
			println(err.Error())
			return
		}

		println("Rows affected: ", rowCount)
	},
}
