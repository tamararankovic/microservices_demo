package handlers

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	ordering "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type OrderingHandler struct {
	clientHost string
	clientPort string
}

func NewOrderingHandler(host, port string) Handler {
	return &OrderingHandler{
		clientHost: host,
		clientPort: port,
	}
}

func (handler *OrderingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/order/{orderId}/details", handler.GetDetails)
	if err != nil {
		panic(err)
	}
}

func (handler *OrderingHandler) newOrderingClient() ordering.OrderingServiceClient {
	address := fmt.Sprintf("%s:%s", handler.clientHost, handler.clientPort)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Ordering service: %v", err)
	}
	return ordering.NewOrderingServiceClient(conn)
}

func (handler *OrderingHandler) GetDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["orderId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("details endpoint, ID=%s", id)))
}
