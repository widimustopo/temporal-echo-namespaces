package main

import (
	"github.com/joho/godotenv"
	//"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"github.com/widimustopo/temporal-namespaces-manager/temporal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	//init temporal client
	temporalClient := initTemporalClient(loadConfig)

	startWorkflow := initStarterWorkflow(temporalClient)

	defer func() {
		startWorkflow.Stop()
	}()

}

func initTemporalClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}

func initStarterWorkflow(temporalClient client.Client) worker.Worker {
	workerOptions := worker.Options{
		MaxConcurrentWorkflowTaskExecutionSize: libs.AxConcurrentWorkflowSize,
	}
	worker := worker.New(temporalClient, libs.StartActivityWorkflow, workerOptions)
	worker.RegisterWorkflow(workflows.StartActivityWorkflow)

	err := worker.Start()
	if err != nil {
		logrus.Fatal("cannot start temporal worker: " + err.Error())
	}

	return worker
}
