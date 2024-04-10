package command

import (
	"fmt"

	"github.com/SachaCR/spanny/internal/pkg/dbops"

	"github.com/spf13/cobra"
)

var createDatabaseCmd = &cobra.Command{
	Use:   "create-db <database name>",
	Short: "Create a Spanner database with the given name or default to the config file value",
	Run: func(cmd *cobra.Command, args []string) {

		projectId := config.ProjectId
		instanceId := config.InstanceId
		databaseId := config.DatabaseId

		if len(args) > 0 {
			databaseId = args[0]
		}

		_, err := dbops.CreateDatabase(instanceId, projectId, databaseId)

		if err != nil {
			println(err.Error())
			return
		}

		fmt.Printf("Database created: %s\n", databaseId)
	},
}
