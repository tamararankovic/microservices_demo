package domain

import (
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

func (service *OrderService) Get(id primitive.ObjectID) (*Order, error) {
	return service.store.Get(id)
}

func (service *OrderService) GetAll() ([]*Order, error) {
	return service.store.GetAll()
}
