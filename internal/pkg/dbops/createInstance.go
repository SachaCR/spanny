package dbops

import (
	"context"
	"fmt"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"google.golang.org/api/iterator"
)

func HasInstance(ctx context.Context, projectId string, instanceId string) (*instancepb.Instance, error) {

	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)

	if err != nil {
		return nil, err
	}

	defer instanceAdmin.Close()

	instanceIterator := instanceAdmin.ListInstances(ctx, &instancepb.ListInstancesRequest{
		Parent: fmt.Sprintf("projects/%s", projectId),
	})

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
			instance = inst
		}
	}

	return instance, nil
}

func CreateInstance(ctx context.Context, projectId string, instanceId string) (*instancepb.Instance, error) {

	existingInstance, err := HasInstance(ctx, projectId, instanceId)
	if err != nil {
		return nil, err
	}

	if existingInstance != nil {
		return existingInstance, nil
	}

	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		return nil, err
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
