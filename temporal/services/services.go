package services

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-echo-namespaces/entities"
	"github.com/widimustopo/temporal-echo-namespaces/libs"
	"github.com/widimustopo/temporal-echo-namespaces/repositories"
	temporalClient "github.com/widimustopo/temporal-echo-namespaces/temporal/client"
	"github.com/widimustopo/temporal-echo-namespaces/temporal/workflows/starter"
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

	var newReq *entities.TemporalMemberRequest

	newReq = &entities.TemporalMemberRequest{
		Times:        time.Now().Local(),
		Data:         req,
		WorkflowName: libs.RegisterWorkflow,
	}

	resp, err := starter.ExecuteRegisterWorkflow(ctx.Request().Context(), clientMember, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s Services) Order(ctx echo.Context, req *entities.Payment) (interface{}, error) {
	clientPayment := temporalClient.InitTemporalPaymentClient(s.Config)
	defer clientPayment.Close()

	var newReq *entities.TemporalOrderRequest

	minute := 1 //duration expired to paid
	endPayment := time.Now().Local().Add(time.Minute * time.Duration(minute))

	newReq = &entities.TemporalOrderRequest{
		Data:         req,
		WorkflowName: libs.ExpiredWorkflow,
		TypesTimes:   "until",
		Times:        endPayment,
		TaskType:     "http",
		Task:         "http://localhost:8080/updates/payments/",
		Interval:     time.Second,
		MaxInterval:  time.Minute * 5,
		Attempt:      10,
		Retry:        10,
	}

	resp, err := starter.ExecuteOrderWorkflow(ctx.Request().Context(), clientPayment, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s Services) Payment(ctx echo.Context, req *entities.Paid) (interface{}, error) {
	clientPayment := temporalClient.InitTemporalPaymentClient(s.Config)
	defer clientPayment.Close()

	var newReq *entities.TemporalPaymentRequest

	newReq = &entities.TemporalPaymentRequest{
		Times:        time.Now().Local(),
		Data:         req,
		WorkflowName: libs.PaymentWorkflow,
	}

	resp, err := starter.ExecutePaymentWorkflow(ctx.Request().Context(), clientPayment, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s Services) PaymentFail(ctx echo.Context, req *entities.Paid) (interface{}, error) {
	clientPayment := temporalClient.InitTemporalPaymentClient(s.Config)
	defer clientPayment.Close()

	var newReq *entities.TemporalPaymentRequest

	newReq = &entities.TemporalPaymentRequest{
		Times:        time.Now().Local(),
		Data:         req,
		WorkflowName: libs.PaymentFailWorkflow,
	}

	resp, err := starter.ExecutePaymentFailWorkflow(ctx.Request().Context(), clientPayment, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s Services) AddProduct(ctx echo.Context, req *entities.Product) (interface{}, error) {
	clientProduct := temporalClient.InitTemporalProductClient(s.Config)
	defer clientProduct.Close()

	var newReq *entities.TemporalProductRequest

	newReq = &entities.TemporalProductRequest{
		Times:        time.Now().Local(),
		Data:         req,
		WorkflowName: libs.PaymentFailWorkflow,
		Attempt:      3,
	}

	resp, err := starter.ExecuteProductWorkflow(ctx.Request().Context(), clientProduct, newReq)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
