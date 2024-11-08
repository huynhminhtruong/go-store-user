package grpc

import (
	"fmt"
	"github.com/huynhminhtruong/go-store-user/src/biz/ports"
	"github.com/huynhminhtruong/go-store-user/src/config"
	"github.com/huynhminhtruong/go-store-user/src/services/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api  ports.APIPort
	port int
	user.UnimplementedUserServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, a)

	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("user-service is running")
}
