package cmd

import "fmt"

func getDatabasePath() string {
	return fmt.Sprintf("projects/%s/instances/%s/databases/%s", config.ProjectId, config.InstanceId, config.DatabaseId)
}
