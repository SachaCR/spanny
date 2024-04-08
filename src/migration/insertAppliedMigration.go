package migration

import (
	"fmt"
	"spanny/src/dbops"
)

func InsertAppliedMigration(databasePath string, migrationName string) error {
	insertQuery := fmt.Sprintf(`INSERT INTO spanner_migrations (name, applied_at) VALUES ('%s', CURRENT_TIMESTAMP)`, migrationName)
	_, err := dbops.ExecuteUpdateQuery(databasePath, insertQuery)
	if err != nil {
		return err
	}

	return nil
}
