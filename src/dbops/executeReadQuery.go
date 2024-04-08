package dbops

import (
	"context"

	"cloud.google.com/go/spanner"
)

func ExecuteReadQuery(databasePath string, query string) (columns []string, rows [][]string, err error) {
	ctx := context.Background()

	client, err := spanner.NewClient(ctx, databasePath)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	stmt := spanner.Statement{SQL: query}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	return rowIterToString(iter)
}
