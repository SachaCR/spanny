package dbops

import (
	"context"

	"cloud.google.com/go/spanner"
)

func UpdateDML(databasePath string, statements []string) (rowCounts []int64, err error) {
	ctx := context.Background()

	client, err := spanner.NewClient(ctx, databasePath)
	if err != nil {
		return []int64{0}, err
	}
	defer client.Close()

	rowCounts = []int64{}

	spannerStatements := make([]spanner.Statement, len(statements))

	for i, statement := range statements {
		spannerStatements[i] = spanner.Statement{
			SQL: statement,
		}

	}

	_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		count, err := txn.BatchUpdate(ctx, spannerStatements)

		if err != nil {
			return err
		}

		rowCounts = count

		return nil
	})

	if err != nil {
		return []int64{0}, err
	}

	return rowCounts, nil
}
