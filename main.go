package main

import (
	"log"
	"online-shopping/internal/server/handlers/account"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 路由器
	router := gin.Default()

	// 初始化 AccountHandler，并将其注册到路由器中
	accountHandler := account.NewAccountHandler()
	accountHandler.RegisterRoutes(router)

	// 启动服务器
	if err := router.Run(":3002"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
