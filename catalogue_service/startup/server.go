package startup

import (
	"fmt"
	"github.com/tamararankovic/microservices_demo/catalogue_service/domain"
	"github.com/tamararankovic/microservices_demo/catalogue_service/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/catalogue_service/infrastructure/persistence"
	"github.com/tamararankovic/microservices_demo/catalogue_service/startup/config"
	catalogue "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
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
	catalogue.RegisterCatalogueServiceServer(grpcServer, controller)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initController() (*api.ProductController, error) {
	service, err := server.initService()
	if err != nil {
		return nil, err
	}
	return api.NewProductController(service), nil
}

func (server *Server) initService() (*domain.ProductService, error) {
	store, err := server.initStore()
	if err != nil {
		return nil, err
	}
	service := domain.NewProductService(store)
	return service, nil
}

func (server *Server) initStore() (domain.ProductStore, error) {
	store, err := persistence.NewProductMongoDBStore(server.config.CatalogueDBHost, server.config.CatalogueDBPort)
	if err != nil {
		return nil, err
	}
	store.DeleteAll()
	for _, product := range products {
		store.Insert(product)
	}
	return store, nil
}
