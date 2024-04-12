package migration

import (
	"context"
	"fmt"

	"github.com/SachaCR/spanny/internal/pkg/dbops"
)

func DeleteRollbackedMigration(ctx context.Context, databasePath string, migrationName string) error {
	insertQuery := fmt.Sprintf(` DELETE FROM spanner_migrations WHERE name = '%s'`, migrationName)
	_, err := dbops.ExecuteUpdateQuery(ctx, databasePath, insertQuery)
	if err != nil {
		return err
	}

	return nil
}
