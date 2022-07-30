package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 使用httptest 测试get post 接口
func main() {
	r := gin.Default()
	r.GET("/get", testGet)
	r.POST("/post", testPost)
	r.Run(":8080")
}

func testGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"Msg":  "这是测试get方法的",
	})
}

func testPost(c *gin.Context) {
	a := new(test)
	if err := c.ShouldBind(&a); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "获取参数失败",
		})
		fmt.Println(a, err)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "这是post测试接口",
		"data": a,
	})
}

type test struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}
