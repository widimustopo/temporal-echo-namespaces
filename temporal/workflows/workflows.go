package workflows

import (
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"go.temporal.io/sdk/workflow"
)

func StartActivityWorkflow(ctx workflow.Context, queueName, req interface{}) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Start Activity Workflow")
		return
	}

	//	ctx = withActivityOptions(ctx, queueName)
	err = workflow.ExecuteActivity(ctx, libs.ActivityRegisterMember, req).Get(ctx, &resp)
	return

}
