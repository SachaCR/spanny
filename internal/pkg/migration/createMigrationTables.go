package migration

import "github.com/SachaCR/spanny/internal/pkg/dbops"

func CreateMigrationTables(databasePath string) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS spanner_migrations (
			name            STRING(MAX) NOT NULL,
			applied_at      TIMESTAMP NOT NULL
		) PRIMARY KEY (name)`,

		`CREATE TABLE IF NOT EXISTS spanner_migrations_lock (
			id              INT64 NOT NULL,
			is_locked       BOOL NOT NULL DEFAULT (FALSE)
		) PRIMARY KEY (id)`,
	}

	return dbops.UpdateDDL(databasePath, statements)
}
