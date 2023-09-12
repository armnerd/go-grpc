package server

import (
	"go-grpc/internal/api"
	"go-grpc/internal/base"

	"go-grpc/internal/data"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewGRPCServer, NewHttpServer, base.Provider, api.Provider, data.Provider)
