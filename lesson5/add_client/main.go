/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-03 11:18:03
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-03 11:46:00
 * @FilePath: \go-micro\lesson5\add_client\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	proto "add_client/pb"
	"context"
	"fmt"
	"log"
	"time"

	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	X int32 = int32(*flag.Int("x", 23, "x"))
	Y int32 = int32(*flag.Int("y", 20, "y"))
)

// add rpc client
func main() {
	//从命令行解析x y的值
	flag.Parse()

	//1 连接rpc server
	conn, err := grpc.Dial("127.0.0.1:8973", grpc.WithTransportCredentials(insecure.NewCredentials())) //需要传一个认证，我们这里使用新建一个
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
	}
	defer conn.Close()

	//2 创建rpc 客户端
	client := proto.NewCalcServiceClient(conn)

	//3 发起rpc调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{X: int32(X), Y: int32(Y)})
	if err != nil {
		log.Fatalf("发生错误 %v", err)
	}

	//4 打印结果
	fmt.Println(resp.GetResult())
}
