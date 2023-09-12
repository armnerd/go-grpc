package server

import (
	"fmt"
	"go-grpc/internal/api"

	"go-grpc/internal/base"

	"github.com/gin-gonic/gin"
)

// http 服务
type HttpServer struct {
	api  api.Api
	conf base.Conf
}

func NewHttpServer(api *api.Api, conf base.Conf) IServerHttp {
	return &HttpServer{
		api:  *api,
		conf: conf,
	}
}

func (server *HttpServer) Start() error {
	conf := server.conf.Server
	app := gin.New()
	app = route(app, server)
	app.Run(":" + conf.PortHttp)
	return nil
}

func (server *HttpServer) Stop() error {
	fmt.Println("http server shutdown")
	return nil
}
