package config

import "os"

type Config struct {
	Port           string
	OrderingDBHost string
	OrderingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("ORDERING_SERVICE_PORT"),
		OrderingDBHost: os.Getenv("ORDERING_DB_HOST"),
		OrderingDBPort: os.Getenv("ORDERING_DB_PORT"),
	}
}
