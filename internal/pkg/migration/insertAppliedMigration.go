package migration

import (
	"context"
	"fmt"

	"github.com/SachaCR/spanny/internal/pkg/dbops"
)

func InsertAppliedMigration(ctx context.Context, databasePath string, migrationName string) error {
	insertQuery := fmt.Sprintf(`INSERT INTO spanner_migrations (name, applied_at) VALUES ('%s', CURRENT_TIMESTAMP)`, migrationName)
	_, err := dbops.ExecuteUpdateQuery(ctx, databasePath, insertQuery)
	if err != nil {
		return err
	}

	return nil
}
