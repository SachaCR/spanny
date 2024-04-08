package cmd

import (
	"context"
	"spanny/src/ui"

	"cloud.google.com/go/spanner"

	"github.com/spf13/cobra"
)

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Displays the current state of migrations",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		databasePath := getDatabasePath()

		client, err := spanner.NewClient(ctx, databasePath)

		if err != nil {
			println(err.Error())
			return
		}

		defer client.Close()

		stmt := spanner.Statement{SQL: `SELECT * FROM spanner_migrations`}
		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()

		ui.RenderTableFromRowIterator(iter)

		stmtLock := spanner.Statement{SQL: `SELECT * FROM spanner_migrations_lock`}
		iterLock := client.Single().Query(ctx, stmtLock)
		defer iterLock.Stop()

		ui.RenderTableFromRowIterator(iterLock)
	},
}
