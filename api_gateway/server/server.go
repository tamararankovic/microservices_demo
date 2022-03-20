package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	config2 "github.com/tamararankovic/microservices_demo/api_gateway/config"
	"github.com/tamararankovic/microservices_demo/api_gateway/handlers"
	catalogueGw "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	orderingGw "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config   *config2.Config
	mux      *runtime.ServeMux
	handlers []handlers.Handler
}

func NewServer(config *config2.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initCustomHandlers()
	return server
}

func (server *Server) initCustomHandlers() {
	orderingHandler := handlers.NewOrderingHandler(server.config.OrderingHost, server.config.OrderingPort)
	orderingHandler.Init(server.mux)
	server.handlers = append([]handlers.Handler{}, orderingHandler)
}

func (server *Server) Start() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	err := catalogueGw.RegisterCatalogueServiceHandlerFromEndpoint(context.TODO(), server.mux, catalogueEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	err = orderingGw.RegisterOrderingServiceHandlerFromEndpoint(context.TODO(), server.mux, orderingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
