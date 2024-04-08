package cmd

import (
	"fmt"
	"spanny/src/dbops"

	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"

	"github.com/spf13/cobra"
)

var createInstanceCmd = &cobra.Command{
	Use:   "create-instance <instance name>",
	Short: "Create a Spanner instance with the given name or default to the config file value",
	Run: func(cmd *cobra.Command, args []string) {
		println(len(args))

		projectId := config.ProjectId
		instanceId := config.InstanceId

		if len(args) > 0 {
			instanceId = args[0]
		}

		instance, err := dbops.CreateInstance(projectId, instanceId)

		if err != nil {
			println(err.Error())
			return
		}

		// The instance may not be ready to serve yet.
		if instance.State != instancepb.Instance_READY {
			println("Instance not ready yet")
			return
		}

		fmt.Printf("Instance created: %s\n", instanceId)
	},
}
