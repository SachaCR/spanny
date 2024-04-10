package command

import (
	"context"
	"fmt"
	"regexp"

	"github.com/SachaCR/spanny/internal/pkg/ui"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var listDatabasesCmd = &cobra.Command{
	Use:   "list-db",
	Short: "List databases",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		projectId := config.ProjectId
		instanceId := config.InstanceId

		databaseAdminClient, err := database.NewDatabaseAdminClient(ctx)
		if err != nil {
			println(err.Error())
			return
		}

		defer databaseAdminClient.Close()

		parent := fmt.Sprintf("projects/%s/instances/%s", projectId, instanceId)

		iter := databaseAdminClient.ListDatabases(ctx, &databasepb.ListDatabasesRequest{
			Parent:    parent,
			PageSize:  100,
			PageToken: "",
		})

		var columns = []string{"Databases"}
		var databases [][]string
		for {
			database, err := iter.Next()

			if err == iterator.Done {
				break
			}

			if err != nil {
				println(err.Error())
				return
			}
			databasePath := database.GetName()

			matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(databasePath)

			databases = append(databases, []string{matches[2]})
		}

		ui.RenderTable(columns, databases)
	},
}
