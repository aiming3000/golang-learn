package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "mall/web/controller/user"
	"net/http"
)

const addr = ":9000"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Gin访问成功",
	})
}

func main() {
	r := gin.Default()
	r.Handle("GET", "/", Index)
	r.Handle("POST", "/test-req", controller.AdminLogin)
	if err := r.Run(addr); err != nil {
		fmt.Println("err")
	}

}
