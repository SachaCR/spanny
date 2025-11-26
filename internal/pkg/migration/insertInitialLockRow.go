package migration

import (
	"context"

	"cloud.google.com/go/spanner"
)

func InsertInitialLockRow(ctx context.Context, databasePath string) error {
	client, err := spanner.NewClient(ctx, databasePath)

	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		insertStatement := spanner.Statement{
			SQL: `INSERT INTO spanner_migrations_lock (id, is_locked) VALUES (1, false)`,
		}

		txn.Update(ctx, insertStatement)
		return nil
	})

	if err != nil {
		errorCode := spanner.ErrCode(err)
		if errorCode.String() == "AlreadyExists" {
			return nil
		}

		return err
	}

	return nil
}
