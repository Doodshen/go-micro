package main

import (
	"context"
	"flag"
	"fmt"
	"hello_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

//gRPC客户端
//调用server端的syahello方法

var name = flag.String("name", "kingshen", "通过-name告诉服务端server是谁")

func main() {
	flag.Parse()
	//连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	defer conn.Close()

	//创建客户端
	c := pb.NewGreeterClient(conn)
	//执行rpc调用并打印收到的数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//发起普通RPC调用
	//带元数据
	//1创建元数据
	md := metadata.Pairs(
		"token", "app-test-kingshen",
	)
	ctx = metadata.NewOutgoingContext(ctx, md) //带上元数据

	//客户端接收headre和tradil 需要在发起调用之前创建
	var header, trailer metadata.MD

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("could not gerrt :%v ", err)
	}
	//拿到响应之前获取header
	fmt.Printf("header:%v\n", header)

	//处理响应
	log.Printf("Greeting :%s", r.GetReply())

	//拿到响应数据之后可以获取trailer
	fmt.Printf("tradiler:%#v\n", trailer)
}
