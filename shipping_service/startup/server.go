package startup

import (
	"fmt"
	shipping "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	"github.com/tamararankovic/microservices_demo/shipping_service/domain"
	"github.com/tamararankovic/microservices_demo/shipping_service/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/shipping_service/infrastructure/persistence"
	"github.com/tamararankovic/microservices_demo/shipping_service/startup/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	server.startGrpcServer()
}

func (server *Server) startGrpcServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	controller, err := server.initController()
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}
	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServiceServer(grpcServer, controller)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initController() (*api.OrderController, error) {
	service, err := server.initService()
	if err != nil {
		return nil, err
	}
	return api.NewOrderController(service), nil
}

func (server *Server) initService() (*domain.OrderService, error) {
	store, err := server.initStore()
	if err != nil {
		return nil, err
	}
	service := domain.NewOrderService(store)
	return service, nil
}

func (server *Server) initStore() (domain.OrderStore, error) {
	store, err := persistence.NewOrderMongoDBStore(server.config.ShippingDBHost, server.config.ShippingDBPort)
	if err != nil {
		return nil, err
	}
	store.DeleteAll()
	for _, Order := range orders {
		store.Insert(Order)
	}
	return store, nil
}
