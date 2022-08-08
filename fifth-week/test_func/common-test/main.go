package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", TestGet)
	router.POST("/test2", TestPost)
	router.Run(":8089")
}

type testP struct {
	Name string
}

func TestGet(c *gin.Context) {
	a := 10
	if a > 0 {
		c.JSON(200, gin.H{
			"body": "success",
		})
	} else {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "get测试失败",
		})
	}
}
func TestPost(c *gin.Context) {
	tt := testP{}
	if err := c.ShouldBind(&tt); err != nil {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "false",
		})
		return
	}
	if tt.Name != "lll" {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "get测试成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "get测试失败",
		})
	}
}
