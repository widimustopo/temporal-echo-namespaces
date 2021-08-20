package starter

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/entities"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"go.temporal.io/sdk/client"
)

func ExecuteRegisterWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalMemberRequest) (resp interface{}, err error) {

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.RegisterWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.RegisterWorkflow, req)
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

func ExecuteOrderWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalOrderRequest) (resp interface{}, err error) {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.OrderWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.OrderWorkflow, req)
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

func ExecutePaymentWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalPaymentRequest) (resp interface{}, err error) {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.PaymentWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.PaymentWorkflow, req)
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

func ExecutePaymentFailWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalPaymentRequest) (resp interface{}, err error) {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.PaymentFailWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.PaymentFailWorkflow, req)
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
