package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/widimustopo/temporal-echo-namespaces/entities"
)

type Repositories interface {
	Register(ctx echo.Context, req *entities.Member) (interface{}, error)
	Order(ctx echo.Context, req *entities.Payment) (interface{}, error)
	Payment(ctx echo.Context, req *entities.Paid) (interface{}, error)
	PaymentFail(ctx echo.Context, req *entities.Paid) (interface{}, error)
	AddProduct(ctx echo.Context, req *entities.Product) (interface{}, error)
}
