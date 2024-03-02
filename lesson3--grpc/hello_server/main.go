/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-01 19:45:19
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 10:02:00
 * @FilePath: \go-micro\lesson3--grpc\hello_server\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"hello_server/pb"

	"net"

	"google.golang.org/grpc"
)

//grpc server

type server struct {
	pb.UnimplementedGreeterServer //嵌套未实现的结构体，已经实现的覆盖默认的
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello" + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	//设置监听路径
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("faild to listen ,err :%v\n", err)
		return
	}

	s := grpc.NewServer() //创建grpc服务器

	//注册服务  在grpc服务端注册服务
	pb.RegisterGreeterServer(s, &server{})

	//启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v", err)
		return
	}
}
