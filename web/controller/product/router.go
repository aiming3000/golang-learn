package product

import (
	"github.com/gin-gonic/gin"
	"mall/web/middle_ware"
)

func Router(router *gin.RouterGroup) {

	router.GET("/get_product_list", middle_ware.JwtTokenValid, GetProductList)
	router.POST("/product_add", middle_ware.JwtTokenValid, ProductAdd)

}
