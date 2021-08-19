package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/widimustopo/temporal-namespaces-manager/entities"
)

type Repositories interface {
	Register(ctx echo.Context, req *entities.Member) (interface{}, error)
	Order(ctx echo.Context, req *entities.Payment) (interface{}, error)
	Payment(ctx echo.Context, paymentID string) (interface{}, error)
}
