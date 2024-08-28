package product

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.GET("/get_product_list", GetProductList)
	router.POST("/product_add", ProductAdd)

	//router.GET("/get_product_list", middle_ware.JwtTokenValid, GetProductList)
	//router.POST("/product_add", middle_ware.JwtTokenValid, ProductAdd)
}
