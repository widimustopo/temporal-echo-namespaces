package starter

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-echo-namespaces/entities"
	"github.com/widimustopo/temporal-echo-namespaces/libs"
	"go.temporal.io/sdk/client"
)

func ExecuteRegisterWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalMemberRequest) (resp interface{}, err error) {

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.RegisterWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.RegisterWorkflow, req)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		logrus.Error(err.Error())
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
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		logrus.Error(err.Error())
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
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		logrus.Error(err.Error())
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
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return workflowResp, err
}

func ExecuteProductWorkflow(ctx context.Context, temporalClient client.Client, req *entities.TemporalProductRequest) (resp interface{}, err error) {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: libs.AddProductWorkflow,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, libs.AddProductWorkflow, req)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var workflowResp interface{}
	err = workflowRun.Get(ctx, &workflowResp)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return workflowResp, err
}
