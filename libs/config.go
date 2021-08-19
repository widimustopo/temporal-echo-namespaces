package libs

import (
	"os"
)

type Config struct {
	ServerPort        string
	PaymentNamespaces string
	MemberNamespaces  string
	ProductNamespaces string
}

// NewEnvConfig function to get info in current config file
// @param name string
// TODO : get env
func NewEnvConfig() (*Config, error) {
	env := &Config{
		ServerPort:        os.Getenv("ServerPort"),
		PaymentNamespaces: os.Getenv("PaymentNamespaces"),
		MemberNamespaces:  os.Getenv("MemberNamespaces"),
		ProductNamespaces: os.Getenv("ProductNamespaces"),
	}
	return env, nil
}
