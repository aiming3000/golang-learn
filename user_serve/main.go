package main

import (
	"mall/user_serve/handler"
	pb "mall/user_serve/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	admin_user "mall/user_serve/proto/admin_user"
)

var (
	service = "user_serve"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterUserserveHandler(srv.Server(), new(handler.Userserve)); err != nil {
		logger.Fatal(err)
	}
	admin_user.RegisterAdminUserHandler(srv.Server(), new(handler.Adminuser))
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
