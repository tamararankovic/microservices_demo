package api

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

type Controller interface {
	Init(mux *runtime.ServeMux)
}
