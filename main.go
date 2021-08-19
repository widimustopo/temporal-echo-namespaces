package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/middleware/validator"
	"github.com/widimustopo/temporal-namespaces-manager/responses"
	domain "github.com/widimustopo/temporal-namespaces-manager/routes"
)

var loadConfig *libs.Config

func main() {
	//load config from .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	loadConfig, err = libs.NewEnvConfig()
	if err != nil {
		panic(err)

	}

	router := echo.New()

	router.Validator = validator.NewValidator()

	//custom response for not matching routes
	echo.NotFoundHandler = func(c echo.Context) error {
		// render 404 custom response
		return responses.NotFound(c, "Not Matching of Any Routes", nil, "Not Found")
	}

	domain.Routes(router, loadConfig)
	err = router.Start(loadConfig.ServerPort)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
}
