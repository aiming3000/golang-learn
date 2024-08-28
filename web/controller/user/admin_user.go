package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	user_serve "mall/user_serve/proto/admin_user"
	"net/http"
	//etcd1
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "index",
	})
}

func AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)

	// etcd 2
	registerReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//service := micro.NewService()
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Registry(registerReg), //注册中心
	)
	service.Init()
	// 创建微服务客户端
	//client := user_serve.NewTestService("test", service.Client())
	client := user_serve.NewAdminUserService("user_serve", service.Client())
	// 调用服务
	//rsp, err := client.Call(c, &testpb.Request{
	//	Name: c.Query("key"),
	//})

	rsp, err := client.AdminUserlogin(c, &user_serve.AdminUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  rsp.Msg,
	})
}
