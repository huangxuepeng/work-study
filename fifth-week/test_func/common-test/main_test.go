package main

import (
	"fmt"
	"testing"

	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
)

// 解析返回的错误类型
type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func init() {
	router := gin.Default()
	router.POST("/test2", TestPost)
	utils.SetRouter(router) //把启动的engine 对象传入到test框架中
}

// 真正的测试单元
func TestMain_TestPost(t *testing.T) {
	param := make(map[string]interface{})
	param["Name"] = "test"
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("POST", "/test2", "json", param, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}
	fmt.Println(resp) // 表示返回的错误为空
}
