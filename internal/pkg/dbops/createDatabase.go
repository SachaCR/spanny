package dbops

import (
	"context"
	"fmt"
	"regexp"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"google.golang.org/api/iterator"
)

func HasDatabase(ctx context.Context, instanceId string, projectId string, databaseId string) (*databasepb.Database, error) {
	databasePath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, databaseId)

	databaseAdminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	defer databaseAdminClient.Close()

	databaseIterator := databaseAdminClient.ListDatabases(ctx, &databasepb.ListDatabasesRequest{
		Parent: fmt.Sprintf("projects/%s/instances/%s", projectId, instanceId),
	})

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
			database = db
		}
	}

	return database, nil
}

func CreateDatabase(ctx context.Context, instanceId string, projectId string, databaseId string) (*databasepb.Database, error) {

	databasePath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectId, instanceId, databaseId)

	existingDatabase, err := HasDatabase(ctx, instanceId, projectId, databaseId)
	if err != nil {
		return nil, err
	}

	if existingDatabase != nil {
		return existingDatabase, nil
	}

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

	op, err := databaseAdminClient.CreateDatabase(ctx, &databasepb.CreateDatabaseRequest{
		Parent:          instancePath,
		CreateStatement: "CREATE DATABASE`" + databaseId + "`",
	})

	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}
