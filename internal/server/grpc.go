package server

import (
	"fmt"
	"go-grpc/internal/api"
	"log"
	"net"

	"go-grpc/internal/base"
	pb "go-grpc/internal/proto"

	"google.golang.org/grpc"
)

// Grpc 服务
type GrpcServer struct {
	api  api.Api
	conf base.Conf
}

func NewGRPCServer(api *api.Api, conf base.Conf) IServerGrpc {
	return &GrpcServer{
		api:  *api,
		conf: conf,
	}
}

func (server *GrpcServer) Start() error {
	conf := server.conf.Server
	lis, err := net.Listen("tcp", conf.Host+":"+conf.PortGrpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBlogServer(s, &server.api)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

func (server *GrpcServer) Stop() error {
	fmt.Println("grpc server shutdown")
	return nil
}
