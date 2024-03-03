/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-03 14:00:04
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 15:40:51
 * @FilePath: \go-micro\lesson6--grpc_stream\hello_client\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"hello_client/pb"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

// callLotsOfReplies()调用服务端的流式
func callLotsOfReplies(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//调用服务器流式rpc
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Println(err)
		return
	}
	//读取响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("stream Recv failed err:%v", err)
			return
		}
		log.Printf("recv :%v", res.GetReply())
	}
}

// callLotsOfGreetings（） 客户端发送流式数据
func callLotsOfGreetings(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//客户端发送流式数据
	//获得流
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Printf("c.LotsOfGreetings(ctx) failed, err:%v\n", err)
		return
	}

	names := []string{"wzs", "kingshen", "doodshen"}
	for _, name := range names {
		stream.Send(&pb.HelloRequest{Name: name}) //发送数据一定要发送定义的结构体
	}

	// 流式发送结束之后要关闭流
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("stream.CloseAndRecv() failed, err:%v\n", err)
		return
	}
	log.Printf("res:%v\n", res.GetReply())

}

// runBidiHello() 双向流
func runBidiHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc //使用<-waitc阻塞，等到通道关闭时结束阻塞
	//协程中循环调用 stream.Recv() 方法，直到服务端关闭连接或者发生错误。如果 stream.Recv() 返回了 io.EOF 错误，说明已经读取完了所有服务端的响应，我们就可以关闭 waitc 通道并返回。在主协程中，我们使用 <-waitc 语法阻塞等待 waitc 通道关闭。
}

func main() {
	//连接server
	conn, err := grpc.Dial("128.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()

	//创建客户端----使用pb
	c := pb.NewGreeterClient(conn)

	runBidiHello(c)

}
