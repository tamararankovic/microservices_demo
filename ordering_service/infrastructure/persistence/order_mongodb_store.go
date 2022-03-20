package persistence

import (
	"context"
	"github.com/tamararankovic/microservices_demo/ordering_service/application"
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "order"
	COLLECTION = "order"
)

type OrderMongoDBStore struct {
	orders *mongo.Collection
}

func NewOrderMongoDBStore(host, port string) (application.OrderStore, error) {
	client, err := GetClient(host, port)
	if err != nil {
		return nil, err
	}
	orders := client.Database(DATABASE).Collection(COLLECTION)
	store := &OrderMongoDBStore{
		orders: orders,
	}
	return store, nil
}

func (store *OrderMongoDBStore) Get(id primitive.ObjectID) (*domain.Order, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *OrderMongoDBStore) GetAll() ([]*domain.Order, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *OrderMongoDBStore) Insert(Order *domain.Order) error {
	result, err := store.orders.InsertOne(context.TODO(), Order)
	if err != nil {
		return err
	}
	Order.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *OrderMongoDBStore) DeleteAll() {
	store.orders.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *OrderMongoDBStore) filter(filter interface{}) ([]*domain.Order, error) {
	cursor, err := store.orders.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *OrderMongoDBStore) filterOne(filter interface{}) (Order *domain.Order, err error) {
	result := store.orders.FindOne(context.TODO(), filter)
	err = result.Decode(&Order)
	return
}

func decode(cursor *mongo.Cursor) (orders []*domain.Order, err error) {
	for cursor.Next(context.TODO()) {
		var Order domain.Order
		err = cursor.Decode(&Order)
		if err != nil {
			return
		}
		orders = append(orders, &Order)
	}
	err = cursor.Err()
	return
}
