package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	user_serve "mall/user_serve/proto/admin_user"
	"mall/web/utils"
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

func AdminLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
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

	rep, err := client.AdminUserlogin(ctx, &user_serve.AdminUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})

	} else {
		admin_token, err1 := utils.GenToken(username, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)
		if err1 != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": rep.Code,
				"msg":  rep.Msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":        rep.Code,
				"msg":         rep.Msg,
				"user_name":   username,
				"admin_token": admin_token,
			})
		}

	}

}
