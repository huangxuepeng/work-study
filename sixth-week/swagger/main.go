package main

import (
	// _ "sixth-week/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/test/:id", test)
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler,
		"NAME_OF_ENV_VARIABLE"))
	r.Run(":8080")
}

// @Summary 测试接口
// @Description 测试函数
// @Tag  测试
// @Router /test/{id} [GET]
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} Response
func test(c *gin.Context) {
	var d IDD
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"Msg":  "失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"Msg":  "测试成功",
	})
}

type IDD struct {
	Id int
}
type Response struct {
	Id   int
	Name string
}
