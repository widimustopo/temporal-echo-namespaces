package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/temporal/controller"
)

func Routes(router *echo.Echo, cfg *libs.Config) {
	serviesHandler := controller.NewServiceHandler(cfg)
	router.POST("/registers", serviesHandler.Register)
	router.POST("/orders/:product_id", serviesHandler.Order)
	router.POST("/payment/:payment_id", serviesHandler.Payment)
}
