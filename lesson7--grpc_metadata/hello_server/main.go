package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"strconv"
	"time"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//grpc server

type server struct {
	pb.UnimplementedGreeterServer //嵌套未实现的结构体，已经实现的覆盖默认的
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	//利用defer使响应数据后，发送trailer
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()

	//在执行业务逻辑之前要check  matadata中是否包含token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}

	//判断元数据中是否有需要的值
	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-kingshen" {
		return nil, status.Error(codes.Unauthenticated, "无效的token")
	}
	//判断完成以后进行有效的业务处理
	reply := "hello" + in.GetName()

	//发送数据前发送header
	header := metadata.New(map[string]string{
		"location": "heze",
	})
	grpc.SendHeader(ctx, header)
	//发送数据
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
