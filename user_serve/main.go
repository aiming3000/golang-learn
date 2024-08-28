package main

import (
	"mall/user_serve/handler"
	pb "mall/user_serve/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	admin_user "mall/user_serve/proto/admin_user"

	//etcd1
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

var (
	service = "user_serve"
	version = "latest"
)

// etcd 2
const (
	ServerName = "etcd-test"
	EtcdAddr   = "127.0.0.1:2379"
)

func main() {
	//registerReg := etcd.NewRegistry()
	//可以设置端口 不设置默认是127.0.0.1:2379
	registerReg := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		//micro.Name(ServerName),      //服务名称
		micro.Registry(registerReg), //注册中心
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
