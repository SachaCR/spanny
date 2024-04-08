package migration

import (
	"fmt"
	"spanny/src/dbops"
)

func DeleteRollbackedMigration(databasePath string, migrationName string) error {
	insertQuery := fmt.Sprintf(` DELETE FROM spanner_migrations WHERE name = '%s'`, migrationName)
	_, err := dbops.ExecuteUpdateQuery(databasePath, insertQuery)
	if err != nil {
		return err
	}

	return nil
}
