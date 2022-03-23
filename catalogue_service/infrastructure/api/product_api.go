package api

import (
	"context"
	"github.com/tamararankovic/microservices_demo/catalogue_service/domain"
	pb "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	pb.UnimplementedCatalogueServiceServer
	service *domain.ProductService
}

func NewProductController(service *domain.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (controller *ProductController) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	product, err := controller.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	productPb := mapProduct(product)
	response := &pb.GetResponse{
		Product: productPb,
	}
	return response, nil
}

func (controller *ProductController) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	products, err := controller.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Products: []*pb.Product{},
	}
	for _, product := range products {
		current := mapProduct(product)
		response.Products = append(response.Products, current)
	}
	return response, nil
}
