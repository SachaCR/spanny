package dbops

import (
	"context"
	"fmt"
	"regexp"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"google.golang.org/api/iterator"
)

func CreateDatabase(instanceId string, projectId string, databaseId string) (*databasepb.Database, error) {
	ctx := context.Background()

	databasePath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, databaseId)

	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(databasePath)
	if matches == nil || len(matches) != 3 {
		error := fmt.Errorf("INVALID DATABASE ID %s", databasePath)
		return nil, error
	}

	instancePath := matches[1]

	databaseAdminClient, err := database.NewDatabaseAdminClient(ctx)

	if err != nil {
		return nil, err
	}

	defer databaseAdminClient.Close()

	databaseIterator := databaseAdminClient.ListDatabases(ctx, &databasepb.ListDatabasesRequest{
		Parent: fmt.Sprintf("projects/%s/instances/%s", projectId, instanceId),
	})

	hasDatabase := false
	var database *databasepb.Database

	for {
		db, err := databaseIterator.Next()

		if err != nil {
			if err == iterator.Done {
				break
			}

			return nil, err
		}

		if db.GetName() == databasePath {
			hasDatabase = true
			database = db
		}
	}

	if hasDatabase {
		return database, nil
	}

	op, err := databaseAdminClient.CreateDatabase(ctx, &databasepb.CreateDatabaseRequest{
		Parent:          instancePath,
		CreateStatement: "CREATE DATABASE`" + databaseId + "`",
	})

	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}
