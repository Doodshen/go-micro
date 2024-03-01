/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-29 11:02:52
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-29 11:27:34
 * @FilePath: \go-micro\lesson1\http_add\server\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//编写基于Restful API的调用方式

// 定义方法请求参数
type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// 定义结果结构体
type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

// 定义远程方法
func add(x, y int) int {
	return x + y
}

// 解析请求并返回响应
func addHandler(w http.ResponseWriter, r *http.Request) {
	//解析参数
	b, _ := ioutil.ReadAll(r.Body)
	var param addParam
	json.Unmarshal(b, &param)

	//业务逻辑
	ret := add(param.X, param.Y)

	//返回响应
	respoBytes, _ := json.Marshal(addResult{Code: 0, Data: ret})
	w.Write(respoBytes)
}

func main() {
	http.HandleFunc("/add", addHandler) //注册路由

	log.Fatal(http.ListenAndServe(":9090", nil)) //启动服务
}
