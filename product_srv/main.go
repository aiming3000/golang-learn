package main

import (
	"mall/product_srv/handler"
	pb "mall/product_srv/proto"
	product_srv "mall/product_srv/proto/product"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"

	//etcd1
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

var (
	service = "product_srv"
	version = "latest"
)

func main() {

	registerReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Registry(registerReg), //注册中心
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterProductsrvHandler(srv.Server(), new(handler.Productsrv)); err != nil {
		logger.Fatal(err)
	}
	//注册商品处理逻辑
	product_srv.RegisterProductsHandler(srv.Server(), new(handler.Product))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
