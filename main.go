// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 创建默认的 Gin 路由器
	router := gin.Default()

	// 创建一个路由组 /group
	group := router.Group("/group")
	{
		group.POST("/hello", HelloHandler)
		group.POST("/calc", CalcHandler)
	}

	// 运行 HTTP 服务，监听 3002 端口
	if err := router.Run(":3002"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
