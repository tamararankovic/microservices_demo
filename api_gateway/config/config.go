package config

import "os"

type Config struct {
	Port          string
	CatalogueHost string
	CataloguePort string
	OrderingHost  string
	OrderingPort  string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("GATEWAY_PORT"),
		CatalogueHost: os.Getenv("CATALOGUE_SERVICE_HOST"),
		CataloguePort: os.Getenv("CATALOGUE_SERVICE_PORT"),
		OrderingHost:  os.Getenv("ORDERING_SERVICE_HOST"),
		OrderingPort:  os.Getenv("ORDERING_SERVICE_PORT"),
	}
}
