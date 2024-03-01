// client/main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//基于RESTful API的服务调用的客户端

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	url := "http://127.0.0.1:9090/add"
	param := addParam{
		X: 1,
		Y: 2,
	}

	paramBytes, _ := json.Marshal(param)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(paramBytes))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)

	fmt.Println(respData.Data)

}
