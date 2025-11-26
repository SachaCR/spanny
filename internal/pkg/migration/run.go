package migration

import (
	"context"
	"os"
	"strings"

	"github.com/SachaCR/spanny/internal/pkg/dbops"
)

type MigrationDirection string

const (
	Up   MigrationDirection = "up"
	Down MigrationDirection = "down"
)

type ApplyMigrationParams struct {
	Direction          MigrationDirection
	DatabasePath       string
	MigrationName      string
	MigrationFilesPath string
}

func RunMigration(ctx context.Context, params ApplyMigrationParams) error {
	databasePath := params.DatabasePath
	migrationName := params.MigrationName
	migrationFilesPath := params.MigrationFilesPath
	direction := params.Direction

	fileName := "up.sql"

	if direction == Down {
		fileName = "down.sql"
	}

	fileBytes, err := os.ReadFile(migrationFilesPath + "/" + migrationName + "/" + fileName)
	if err != nil {
		return err
	}

	fileContent := string(fileBytes)

	migrationStatements := strings.Split(fileContent, ";")

	if strings.Contains(migrationName, "_DML_") {
		_, err = dbops.UpdateDML(ctx, databasePath, migrationStatements[0:len(migrationStatements)-1])
		if err != nil {
			return err
		}
	}

	if strings.Contains(migrationName, "_DDL_") {
		err = dbops.UpdateDDL(ctx, databasePath, migrationStatements[0:len(migrationStatements)-1])
		if err != nil {
			return err
		}
	}

	if direction == Down {
		err = DeleteRollbackedMigration(ctx, databasePath, migrationName)
		if err != nil {
			return err
		}

		return nil
	}

	err = InsertAppliedMigration(ctx, databasePath, migrationName)
	if err != nil {
		return err
	}

	return nil
}
