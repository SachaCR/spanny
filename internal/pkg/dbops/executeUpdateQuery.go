package dbops

import (
	"context"

	"cloud.google.com/go/spanner"
)

func ExecuteUpdateQuery(ctx context.Context, databasePath string, updateQuery string) (rowCount int64, err error) {

	client, err := spanner.NewClient(ctx, databasePath)
	if err != nil {
		println(err.Error())
		return
	}
	defer client.Close()

	rowCount = int64(0)

	_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		updateStatement := spanner.Statement{
			SQL: updateQuery,
		}

		count, err := txn.Update(ctx, updateStatement)

		if err != nil {
			return err
		}

		rowCount = count

		return nil
	})

	if err != nil {
		return 0, err
	}

	return rowCount, nil
}
