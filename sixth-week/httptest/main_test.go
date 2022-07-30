package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   test   `json:"data"`
}

func config() *gin.Engine {
	router := gin.Default()
	router.GET("/get", testGet)
	router.POST("/post", testPost)
	// router.Run(":8080")
	return router
}

func Get(url string, router *gin.Engine) []byte {

	// 构造get请求
	req := httptest.NewRequest("GET", url, nil)

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func Post(uri string, param url.Values, router *gin.Engine) []byte {
	// 构造post 请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handdler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
func TestTestGet(t *testing.T) {
	router := config()
	// 初始化请求地址
	url := "/get"
	body := Get(url, router)
	fmt.Println(string(body))
}

func TestTestPost(t *testing.T) {
	router := config()
	uri := "/post"
	params := url.Values{
		"name": {"test"},
		"age":  {"12"},
	}
	// 发起post请求，以表单形式传递参数
	body := Post(uri, params, router)
	//body := PostForm3(uri, param, router)

	// 解析响应，判断响应是否与预期一致
	response := &Response{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	t.Log(response.Data)

}
