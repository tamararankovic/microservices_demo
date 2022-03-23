package config

import "os"

type Config struct {
	Port           string
	ShippingDBHost string
	ShippingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("SHIPPING_SERVICE_PORT"),
		ShippingDBHost: os.Getenv("SHIPPING_DB_HOST"),
		ShippingDBPort: os.Getenv("SHIPPING_DB_PORT"),
	}
}
