package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	store ProductStore
}

func NewProductService(store ProductStore) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (service *ProductService) Get(id primitive.ObjectID) (*Product, error) {
	return service.store.Get(id)
}

func (service *ProductService) GetAll() ([]*Product, error) {
	return service.store.GetAll()
}
