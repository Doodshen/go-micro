package main

import (
	"boostore/pb"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//1 定义grpc service：proto文件中定义
//2 实现 gRPC service ：就是bookstore.go中实现的
//2 创建gRPC Server：创建一个 gRPC Server 对象，用于接收和处理客户端发送过来的请求。在创建 Server 时，我们需要指定要监听的网络地址和端口号，并将 service 注册到 Server 上。
//3 注册
//3 启动gRPC Server服务

func main() {
	//连接数据库
	dsn := "root:abc123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := NewDB(dsn)
	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
		return
	}

	//创建grpc Server
	srv := server{
		bs: &bookstore{db: db},
	}
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}
	s := grpc.NewServer()

	//注册服务
	pb.RegisterBookstoreServer(s, &srv)

	go func() {
		fmt.Println(s.Serve(lis))
	}()

	//grpc——Geteway
	//1.创建 gRPC 连接
	//2 创建 ServeMux：
	//3 创建 HTTP Server：
	//4 启动 HTTP Server

	//1 创建grpc连接
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(), //阻塞 直到连接上grpc服务
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("grpc conn failed err:%v", err)
	}

	//2 创建ServeMux:
	gwmux := runtime.NewServeMux()
	pb.RegisterBookstoreHandler(context.Background(), gwmux, conn)

	//3创建http server
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	//4 启动服务
	fmt.Println("grpc_Getway serve on :8080")
	gwServer.ListenAndServe()

}
