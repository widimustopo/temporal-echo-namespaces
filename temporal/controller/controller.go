package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/widimustopo/temporal-namespaces-manager/entities"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/repositories"
	"github.com/widimustopo/temporal-namespaces-manager/responses"
	"github.com/widimustopo/temporal-namespaces-manager/temporal/services"
)

type ServicesHandler struct {
	repository repositories.Repositories
	config     *libs.Config
}

func NewServiceHandler(cfg *libs.Config) *ServicesHandler {

	repo := services.NewServices(cfg)

	return &ServicesHandler{
		repository: repo,
		config:     cfg,
	}
}

func (s ServicesHandler) Register(ctx echo.Context) error {
	var req *entities.Member
	if err := ctx.Bind(&req); err != nil {
		return responses.BadRequest(ctx, libs.BadRequest, nil, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return responses.ValidationError(ctx, libs.ValidationError, nil, err.Error())
	}

	data, err := s.repository.Register(ctx, req)
	if err != nil {
		return responses.InternalServerError(ctx, libs.InternalServerError, nil, err.Error())
	}

	return responses.SingleData(ctx, libs.OK, data, nil)
}

func (s ServicesHandler) Order(ctx echo.Context) error {
	var req *entities.Payment

	if err := ctx.Bind(&req); err != nil {
		return responses.BadRequest(ctx, libs.BadRequest, nil, err.Error())
	}

	if err := ctx.Validate(req); err != nil {
		return responses.ValidationError(ctx, libs.ValidationError, nil, err.Error())
	}

	data, err := s.repository.Order(ctx, req)
	if err != nil {
		return responses.InternalServerError(ctx, libs.InternalServerError, nil, err.Error())
	}

	return responses.SingleData(ctx, libs.OK, data, nil)
}

func (s ServicesHandler) Payment(ctx echo.Context) error {
	paymentID := ctx.Param("payment_id")

	data, err := s.repository.Payment(ctx, paymentID)
	if err != nil {
		return responses.InternalServerError(ctx, libs.InternalServerError, nil, err.Error())
	}

	return responses.SingleData(ctx, libs.OK, data, nil)
}
