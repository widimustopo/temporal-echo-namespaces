package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/temporal/controller"
)

func Routes(router *echo.Echo, cfg *libs.Config) {
	servicesHandler := controller.NewServiceHandler(cfg)
	router.POST("/registers", servicesHandler.Register)
	router.POST("/orders/:member_id/:product_id", servicesHandler.Order)
	router.POST("/payments/:payment_id/members/:member_id", servicesHandler.Payment)
	router.PATCH("/updates/payments/:payment_id/:member_id", servicesHandler.PaymentFail)
}
