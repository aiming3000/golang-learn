package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"mall/web/utils"
	"net/http"
	"strconv"
	"time"

	//etcd1
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"

	product_srv "mall/product_srv/proto/product"
)

func GetProductList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	fmt.Println(currentPage, pageSize)

	// etcd 2
	registerReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//grpc通信
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Registry(registerReg), //注册中心
	)
	service.Init()
	product_service := product_srv.NewProductsService("product_srv", service.Client())
	rep, err := product_service.ProductList(ctx, &product_srv.ProductsRequest{
		CurrentPage: utils.StrToInt(currentPage),
		PageSize:    utils.StrToInt(pageSize),
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         rep.Code,
			"msg":          rep.Msg,
			"products":     rep.Products,
			"total":        rep.Total,
			"current_page": rep.Current,
			"page_size":    rep.PageSize,
		})
	}

}

func ProductAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	price := ctx.PostForm("price")
	num := ctx.PostForm("num")
	unit := ctx.PostForm("unit")
	pic := ctx.PostForm("pic")
	desc := ctx.PostForm("desc")
	file, err := ctx.FormFile("pic")
	fmt.Println(name, price, num, unit, pic, desc)
	//form, _ := ctx.MultipartForm()
	//fmt.Println("ctx.Request.MultipartForm", form)
	var file_path string
	file_path = ""
	if err != nil {
		fmt.Println(err)
	} else {
		unix_int64 := time.Now().Unix()
		unix_str := strconv.FormatInt(unix_int64, 10)
		file_path = "upload/" + unix_str + file.Filename
		ctx.SaveUploadedFile(file, file_path)
	}

	// etcd 2
	registerReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//grpc通信
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Registry(registerReg), //注册中心
	)
	service.Init()
	product_service := product_srv.NewProductsService("product_srv", service.Client())

	rep, err := product_service.ProductAdd(ctx, &product_srv.ProductAddRequest{
		Name: name,
		//Price: float32(utils.StrToInt(price)),
		Price: utils.StrToFloat32(price),
		Num:   utils.StrToInt(num),
		Unit:  unit,
		Pic:   file_path,
		Desc:  desc,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	}

}
