package application

import (
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{
		store: store,
	}
}

func (service *OrderService) Get(id primitive.ObjectID) (*domain.Order, error) {
	return service.store.Get(id)
}

func (service *OrderService) GetAll() ([]*domain.Order, error) {
	return service.store.GetAll()
}
