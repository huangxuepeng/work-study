package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 实现构建项目的,并且完成是数据值的返回
func main() {
	r := gin.Default()
	r.GET("/", test)
	r.Run(":8081")
}

func test(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"code":    200,
		"message": "test api",
	})
}
