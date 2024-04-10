package migration

import (
	"os"
)

func ListFiles(migrationFolderPath string) ([]string, error) {
	// Ouvrir le dossier
	dir, err := os.Open(migrationFolderPath)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			files = append(files, fileInfo.Name())
		}
	}

	return files, nil
}
