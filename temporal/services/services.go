package services

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/entities"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/repositories"
	temporalClient "github.com/widimustopo/temporal-namespaces-manager/temporal/client"
	"github.com/widimustopo/temporal-namespaces-manager/temporal/workflows"
)

type Services struct {
	*libs.Config
}

func NewServices(cfg *libs.Config) repositories.Repositories {
	return Services{cfg}
}

func (s Services) Register(ctx echo.Context, req *entities.Member) (interface{}, error) {
	clientMember := temporalClient.InitTemporalMemberClient(s.Config)

	defer clientMember.Close()

	var newReq *entities.TemporalRequest

	newReq = &entities.TemporalRequest{
		Times:        time.Now().Local(),
		Data:         req,
		WorkflowName: libs.RegisterWorkflow,
	}

	resp, err := workflows.ExecuteRegisterWorkflow(ctx.Request().Context(), clientMember, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s Services) Order(ctx echo.Context, req *entities.Payment) (interface{}, error) {
	return nil, nil
}
func (s Services) Payment(ctx echo.Context, paymentID string) (interface{}, error) {
	return nil, nil
}
