package main

import (
	cfg "github.com/tamararankovic/microservices_demo/ordering_service/infrastructure/config"
	"github.com/tamararankovic/microservices_demo/ordering_service/startup"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
