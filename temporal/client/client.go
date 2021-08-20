package client

import (
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-manager/libs"
	"go.temporal.io/sdk/client"
)

func InitTemporalMemberClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.MemberNamespaces,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}

func InitTemporalPaymentClient(cfg *libs.Config) client.Client {
	c, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.PaymentNamespaces,
	})

	if err != nil {
		logrus.Fatalln("Unable to create client", err)
	}

	return c
}
