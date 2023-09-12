//go:build wireinject
// +build wireinject

package main

import (
	"go-grpc/internal/server"

	"github.com/google/wire"
)

func NewApp() (*App, func(), error) {
	panic(wire.Build(server.Provider, wire.Struct(new(App), "*")))
}
