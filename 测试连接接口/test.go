//基本的GET请求
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	data.Add("name", "黄雪朋")
	data.Add("age", "18")
	resp, err := http.PostForm("http://127.0.0.1:8080/test/post", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // 读出其中的数据
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}
