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

// @Summary 获取日志列表
// @Description 获取用户所有的操作日志
// @Tags 操作审计
// @accept json
// @Produce json
// @Param page path int true "页码"
// @Param limit path int true "每页多少条信息"
// @Param userId path int true "用户id"
// @Param mod query string false "模块 SrcType"
// @Param type path int true "操作 Type"
// @Success 200 {object} ListResponse
// @Router //api/v1/log/{page}/{limit}/{userId}/{mod}/{type} [get]
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

// 列表数据响应内容
type ListResponse struct {
	Code int
	List []*LogRecord
}

type LogRecord struct {
	Uid          int
	LogType      int32
	SrcType      int32
	UserId       string
	DataName     string
	OperatorTime int64
	CreateTime   int64
}
