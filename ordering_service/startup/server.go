package startup

import (
	"fmt"
	ordering "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	"github.com/tamararankovic/microservices_demo/ordering_service/application"
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"github.com/tamararankovic/microservices_demo/ordering_service/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/ordering_service/infrastructure/config"
	"github.com/tamararankovic/microservices_demo/ordering_service/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
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
	ordering.RegisterOrderingServiceServer(grpcServer, controller)
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

func (server *Server) initService() (*application.OrderService, error) {
	store, err := server.initStore()
	if err != nil {
		return nil, err
	}
	service := application.NewOrderService(store)
	return service, nil
}

func (server *Server) initStore() (application.OrderStore, error) {
	store, err := persistence.NewOrderMongoDBStore(server.config.OrderingDBHost, server.config.OrderingDBPort)
	if err != nil {
		return nil, err
	}
	store.DeleteAll()
	for _, Order := range orders {
		store.Insert(Order)
	}
	return store, nil
}

var orders = []*domain.Order{
	{
		Id:        primitive.NewObjectID(),
		Status:    domain.Approved,
		CreatedAt: time.Now(),
		Items: []domain.OrderItem{
			{
				Product: domain.Product{
					Id:    primitive.NewObjectID().Hex(),
					Color: domain.Color{Code: "R"},
				},
				Quantity: 5,
			},
			{
				Product: domain.Product{
					Id:    primitive.NewObjectID().Hex(),
					Color: domain.Color{Code: "G"},
				},
				Quantity: 3,
			},
		},
	},
}
