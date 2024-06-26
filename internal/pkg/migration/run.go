package migration

import (
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

func RunMigration(params ApplyMigrationParams) error {
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
		_, err = dbops.UpdateDML(databasePath, migrationStatements[0:len(migrationStatements)-1])
		if err != nil {
			return err
		}
	}

	if strings.Contains(migrationName, "_DDL_") {
		err = dbops.UpdateDDL(databasePath, migrationStatements[0:len(migrationStatements)-1])
		if err != nil {
			return err
		}
	}

	if direction == Down {
		err = DeleteRollbackedMigration(databasePath, migrationName)
		if err != nil {
			return err
		}

		return nil
	}

	err = InsertAppliedMigration(databasePath, migrationName)
	if err != nil {
		return err
	}

	return nil
}
