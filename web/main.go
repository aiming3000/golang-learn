package main

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
	"log"
	all_router "mall/web/router"
	"net/http"
	//etcd1
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

const addr = ":9000"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Gin访问成功",
	})
}

func main() {
	//r := gin.Default()
	//r.Handle("GET", "/", Index)
	//r.Handle("POST", "/test-req", controller.AdminLogin)
	//if err := r.Run(addr); err != nil {
	//	fmt.Println("err")
	//}

	r := gin.Default()

	//设置注册中心
	registerReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//注册路由组
	all_router.InitRouter(r)
	service := web.NewService(
		web.Name("web"),
		web.Version("latest"),
		web.Handler(r),
		web.Address(":8081"),
		web.Registry(registerReg),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
