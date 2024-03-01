/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-02-29 13:46:35
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-02-29 13:57:13
 * @FilePath: \go-micro\lesson1\rec_add\client\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//当服务端注册服务以后，客户端能看到一个拥有add方法的ServiceA服务

type Args struct {
	X, Y int
}

func mian() {
	//建立http连接 Dial建立网络连接
	//client, err := rpc.DialHTTP("tcp","127.0.0.1:9091")  //基于http

	//基于tcp协议
	client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing", err)
	}

	//同步调用
	args := &Args{10, 20}
	var reply int

	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("err")
	}

	fmt.Println(reply)

	//异步调用
	var reply2 int
	divcall := client.Go("ServiceA.Add", args, &reply2, nil)
	replycall := <-divcall.Done //接收调用结果
	fmt.Println(replycall.Error)
	fmt.Println(reply2)

}
