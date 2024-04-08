package migration

import "spanny/src/dbops"

func GetLastMigrationApplied(databasePath string) (string, error) {

	_, rows, err := dbops.ExecuteReadQuery(databasePath, `
	SELECT * 
	FROM spanner_migrations 
	ORDER BY applied_at 
	DESC LIMIT 1
	`)

	if err != nil {
		return "", err
	}

	if len(rows) == 0 {
		return "NONE", nil
	}

	return rows[0][0], nil
}
