package main

import (
	"github.com/huynhminhtruong/go-store-user/src/biz/adapters/db"
	"github.com/huynhminhtruong/go-store-user/src/biz/adapters/grpc"
	"github.com/huynhminhtruong/go-store-user/src/biz/application/core/api"
	"github.com/huynhminhtruong/go-store-user/src/config"
	"log"
)

/*
	Hexagonal architecture(Ports and Adapters):

		1. primary adapter(chủ động điều khiển các luồng xử lý của hệ thống từ bên ngoài vào):
			storing application

		2. second adapter(chủ động gửi request ra ngoài để gọi các dịch vụ bên ngoài):
			database
			shipping service
*/

func main() {
	// init database adapter server
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	//	init grpc adapter server
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
