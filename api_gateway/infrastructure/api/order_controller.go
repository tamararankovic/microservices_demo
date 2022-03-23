package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	catalogue "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	ordering "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	shipping "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	"net/http"
)

type OrderingController struct {
	orderingClientAddress  string
	catalogueClientAddress string
	shippingClientAddress  string
}

func NewOrderingController(orderingClientAddress, catalogueClientAddress, shippingClientAddress string) Controller {
	return &OrderingController{
		orderingClientAddress:  orderingClientAddress,
		catalogueClientAddress: catalogueClientAddress,
		shippingClientAddress:  shippingClientAddress,
	}
}

func (controller *OrderingController) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/order/{orderId}/details", controller.GetDetails)
	if err != nil {
		panic(err)
	}
}

func (controller *OrderingController) GetDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["orderId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	orderingClient := services.NewOrderingClient(controller.orderingClientAddress)
	orderInfo, err := orderingClient.Get(context.TODO(), &ordering.GetRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	orderDetails := domain.OrderDetails{
		Id:         orderInfo.Order.Id,
		CreatedAt:  orderInfo.Order.CreatedAt.AsTime(),
		Status:     orderInfo.Order.Status.String(),
		OrderItems: make([]domain.OrderItem, 0),
	}
	shippingClient := services.NewShippingClient(controller.shippingClientAddress)
	shippingInfo, err := shippingClient.Get(context.TODO(), &shipping.GetRequest{Id: id})
	if err == nil {
		orderDetails.ShippingStatus = shippingInfo.Order.Status.String()
		orderDetails.ShippingAddress = shippingInfo.Order.ShippingAddress
	}
	for _, item := range orderInfo.Order.Items {
		itemDetails := domain.OrderItem{
			Quantity: uint16(item.Quantity),
		}
		catalogueClient := services.NewCatalogueClient(controller.catalogueClientAddress)
		catalogueInfo, err := catalogueClient.Get(context.TODO(), &catalogue.GetRequest{Id: item.Product.Id})
		if err == nil {
			itemDetails.Product = domain.Product{
				Name:  catalogueInfo.Product.Name,
				Color: getColorName(item.Product.Color.Code, catalogueInfo.Product.Colors),
			}
		}
		orderDetails.OrderItems = append(orderDetails.OrderItems, itemDetails)
	}
	response, err := json.Marshal(orderDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getColorName(code string, colors []*catalogue.Color) string {
	for _, color := range colors {
		if color.Code == code {
			return color.Name
		}
	}
	return ""
}
