package dbops

import (
	"context"
	"fmt"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"google.golang.org/api/iterator"
)

func CreateInstance(projectId string, instanceId string) (*instancepb.Instance, error) {
	ctx := context.Background()

	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)

	if err != nil {
		return nil, err
	}

	defer instanceAdmin.Close()

	instanceIterator := instanceAdmin.ListInstances(ctx, &instancepb.ListInstancesRequest{
		Parent: fmt.Sprintf("projects/%s", projectId),
	})

	hasInstance := false
	var instance *instancepb.Instance

	for {
		inst, err := instanceIterator.Next()

		if err != nil {
			if err == iterator.Done {
				break
			}

			return nil, err
		}

		if inst.DisplayName == instanceId {
			hasInstance = true
			instance = inst
		}
	}

	if hasInstance {
		return instance, nil
	}

	op, err := instanceAdmin.CreateInstance(ctx, &instancepb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s", projectId),
		InstanceId: instanceId,
		Instance: &instancepb.Instance{
			Config:      fmt.Sprintf("projects/%s/instanceConfigs/%s", projectId, "regional-us-central1"),
			DisplayName: instanceId,
			NodeCount:   1,
			Labels:      map[string]string{"cloud_spanner_samples": "true"},
		},
	})

	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}
