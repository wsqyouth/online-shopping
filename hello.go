// hello.go
package main

/*
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由器
	router := gin.Default()

	// 设置 POST 路由，接受 JSON 请求
	router.POST("/hello", func(c *gin.Context) {
		// 定义请求 JSON 结构体
		var json struct {
			Username string `json:"username" binding:"required"`
		}

		// 绑定 JSON 并检测错误
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// 返回 "hello + username"
		c.JSON(http.StatusOK, gin.H{"message": "hello " + json.Username})
	})

	// 运行 HTTP 服务，监听 3002 端口
	router.Run(":3002")
}
*/
