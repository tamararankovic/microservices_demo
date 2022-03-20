package application

import (
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStore interface {
	Get(id primitive.ObjectID) (*domain.Order, error)
	GetAll() ([]*domain.Order, error)
	Insert(product *domain.Order) error
	DeleteAll()
}
