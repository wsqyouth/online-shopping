// handler.go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HelloHandler 处理 /group/hello 路由
func HelloHandler(c *gin.Context) {
	var json struct {
		Username string `json:"username" binding:"required"`
	}

	// 绑定 JSON 请求数据到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	// 返回 "hello + username"
	c.JSON(http.StatusOK, gin.H{"message": "hello " + json.Username})
}

// CalcHandler 处理 /group/calc 路由
func CalcHandler(c *gin.Context) {
	var json struct {
		A int `json:"a" binding:"required"`
		B int `json:"b" binding:"required"`
	}

	// 绑定 JSON 请求数据到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields 'a' and 'b' are required"})
		return
	}

	// 返回 a 和 b 的和
	result := json.A + json.B
	c.JSON(http.StatusOK, gin.H{"result": result})
}
