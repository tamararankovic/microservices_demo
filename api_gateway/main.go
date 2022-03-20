package main

import (
	"github.com/tamararankovic/microservices_demo/api_gateway/config"
	"github.com/tamararankovic/microservices_demo/api_gateway/server"
)

func main() {
	config := config.NewConfig()
	server := server.NewServer(config)
	server.Start()
}
