package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type Single struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func SingleData(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusOK, Single{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func NotFound(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusNotFound, Single{
		Meta: Meta{
			Code:    http.StatusNotFound,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func BadRequest(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusBadRequest, Single{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func ValidationError(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusUnprocessableEntity, Single{
		Meta: Meta{
			Code:    http.StatusUnprocessableEntity,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func InternalServerError(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusInternalServerError, Single{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}

func ListData(c echo.Context, message string, data interface{}, error interface{}) error {
	return c.JSON(http.StatusOK, Single{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
			Error:   error,
		},
		Data: data,
	})
}
