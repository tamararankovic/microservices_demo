package application

import (
	"github.com/tamararankovic/microservices_demo/catalogue_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductStore interface {
	Get(id primitive.ObjectID) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Insert(product *domain.Product) error
	DeleteAll()
}
