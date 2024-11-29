package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 实例
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// 启动服务
	r.Run(":8080")
}
