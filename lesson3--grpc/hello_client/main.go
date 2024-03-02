package main

import (
	"context"
	"flag"
	"hello_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not gerrt :%v ", err)
	}

	//处理响应
	log.Printf("Greeting :%s", r.GetReply())

}
