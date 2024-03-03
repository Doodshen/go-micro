/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-03 11:00:23
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 11:29:46
 * @FilePath: \go-micro\lesson5\add_server\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	proto "add_server/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedCalcServiceServer
}

func (s server) Add(c context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	sum := in.GetX() + in.GetY()
	return &proto.AddResponse{Result: int64(sum)}, nil
}

func main() {
	//设置监听器
	l, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("net.listen failed err :%v", err)
	}

	//创建一个新的 gRPC 服务器实例。
	s := grpc.NewServer()

	//注册
	proto.RegisterCalcServiceServer(s, &server{})

	//启动服务
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.serve failed err %v", err)
	}
}
