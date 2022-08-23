package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	r := gin.Default()

	r.GET("/test/get", testGet)
	r.POST("/test/post", testPost)
	r.Run(":8080")
}

func testGet(c *gin.Context) {
	// 显示一个get请求返回结果
	type test struct {
		Name string
		Age  int
	}
	res := &test{
		Name: "test",
		Age:  18,
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  res,
	})

}
func testPost(c *gin.Context) {
	// 测试一个post请求产出的数据
	res := &test{}
	if err := c.ShouldBind(&res); err != nil {
		glog.Errorf("获取前端参数失败")
		c.JSON(400, gin.H{
			"code": 400,
			"err":  errors.New("获取前端的参数失败"),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

type test struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}
