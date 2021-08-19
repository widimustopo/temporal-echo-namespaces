package workflows

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/entities"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"go.temporal.io/sdk/client"
)

func ExecuteRegisterWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalRequest) (resp interface{}, err error) {

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.RegisterWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, "RegisterWorkflow", req)
	if err != nil {
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		return nil, err
	}

	return workflowResp, err

}
