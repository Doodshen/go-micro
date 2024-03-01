# 介绍 

grpc是一种高性能rpc框架，能运行在任何环境中 

grpc中，客户端可以像调用本地方法一样调用其他机器上的服务端应用程序的方法，更容易创建分布式应用程序和服务 

grpc是基于定义一个服务，指定可以远程调用的带有参数和返回类型的方法，在服务端程序中实现这个接口并且运行grpc服务端处理客户端调用，在客户端中，有一个stub提供和服务端相同的方法 



## 安装grpc

1. 获取grpc作为项目依赖 

~~~GO
go get google.golang.org/grpc@latest
~~~

2. 安装protocol Buffers

## 安装插件

因为本文我们是使用Go语言做开发，接下来执行下面的命令安装`protoc`的Go插件：

安装go语言插件：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

该插件会根据`.proto`文件生成一个后缀为`.pb.go`的文件，包含所有`.proto`文件中定义的类型及其序列化方法。

安装grpc插件：

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

该插件会生成一个后缀为`_grpc.pb.go`的文件，其中包含：

- 一种接口类型(或存根) ，供客户端调用的服务方法。
- 服务器要实现的接口类型。



## 解析

当我们使用 Protocol Buffers 定义 API 接口时，需要定义 `.proto` 文件，并通过 `protoc` 工具编译生成相应的代码。

例如，假设我们有一个用户服务的 API 接口，定义如下：

```
复制代码syntax = "proto3";

package user;

message GetUserRequest {
  int64 id = 1;
}

message User {
  int64 id = 1;
  string name = 2;
  string email = 3;
}

service UserService {
  rpc GetUser(GetUserRequest) returns (User);
}
```

当我们使用 `protoc` 工具编译该文件时，可以指定 `--go-grpc_out` 参数来生成 gRPC 相关的代码。例如：

```
复制代码$ protoc --go_out=. --go-grpc_out=. user.proto
```

上述命令会生成两个文件：`user.pb.go` 和 `user_grpc.pb.go`。其中

* `user.pb.go` 包含了序列化和反序列化等基本的 Protocol Buffers 功能
* `user_grpc.pb.go` 则包含了 gRPC 服务端和客户端的代码。
* 我们可以直接引用 `_grpc.pb.go` 中的代码来实现 gRPC 服务端和客户端的逻辑。









# grpc开发方式 

grpc开发分为三个步骤：

## 一：编写.proto文件定义服务 

grpc基于定义服务的思想，指定可以通过参数和返回类型远程调用的方法，默认情况下，grpc使用`proto buffers `作为接口定义语言（IDL）来描述服务接口和有效负载信息的结构，可以根据需要使用其他IDL代替 

使用protocol buffers 定义一个HelloService服务 

~~~go
service HelloService{
    rpc SayHello（HelloRequest）returns（HelloResponse）
}

meaage HelloRequest{
    string greeting = 1 ；
}

message HelloResponse{
    string reply = 1 ；
}
~~~

### 四种类型

**在grpc中可以定义四种类型的服务方法**

1. 普通rpc ：客户端向服务端发送一个请求，得到一个响应，像普通函数调用一样 

   ~~~protobuf
   rpc SayHello(HelloRequest) returns (HelloResponse)
   ~~~

2. 服务流式rpc，其中客户端向服务器发送请求，并获得一个流来读取一系列消息，客户端从返回中的流读取，知道没有更多的消息，grpc保证在单个rpc调用中的消息是有序的 

   ~~~~protobuf
   rpc LOtsOfReplies(HelloRequest) returns （stream HelloResponse）
   ~~~~

3. 客户端流式rpc，其中客户端写入一些列消息并将其发送给服务器，同样使用提供的流，一旦客户端完成了消息的写入，它就等待服务器读取消息并返回响应。同样，gRPC 保证在单个 RPC 调用中对消息进行排序。

   ~~~protobuf
   rpc LOtsOfGrertings(stream HelloRequest) returns (HelloResponse)
   ~~~

4. 双向流式rpc，双方使用读写流发送一系列消息。这两个流独立运行，因此客户端和服务器可以按照自己喜欢的顺序读写: 例如，服务器可以等待接收所有客户端消息后再写响应，或者可以交替读取消息然后写入消息，或者其他读写组合。每个流中的消息是有序的。



## 二 ：生成指定语言代码 

在`.proto`文件中的定义好服务之后，grpc提供了生成客户端和服务端代码的protocol buffers 编译器插件 ，使用插件色恒诚需要的语言代码，**通常会在客户端调用这些API，并在服务端实现相应的API**

* 服务端：**服务器实现服务声明的方法，并运行一个`grpc`服务器来处理客户端发来的调用请求，**gRPC 底层会对传入的请求进行解码，执行被调用的服务方法，并对服务响应进行编码。
* 客户端：客户端有一个成为**存根（stub）的本地对象，实现了与服务相同的方法**，然后，客户端可以在本地对象上调用这些方法，将调用的参数包装在适当的 protocol buffers 消息类型中—— gRPC 在向服务器发送请求并返回服务器的 protocol buffers 响应之后进行处理。

## 三：编写业务逻辑代码 

grpc帮我们解决了rpc中的服务调用，数据传输，以及消息解码，剩下的工作就是编写业务逻辑代码 

**在服务端编写业务代码实现具体的服务方法，在客户端按需调用这些方法** 







# grpc入门开发 

## 编写proto代码 

`Protocol Buffers`是一种与语言无关，平台无关的可扩展机制，用于序列化结构化数据。使用`Protocol Buffers`可以一次定义结构化的数据，然后可以使用特殊生成的源代码轻松地在各种数据流中使用各种语言编写和读取结构化数据

~~~go
syntax = "proto3"; // 版本声明，使用Protocol Buffers v3版本

option go_package = "xx";  // 指定生成的Go代码在你项目中的导入路径

package pb; // 包名


// 定义服务
service Greeter {
    // SayHello 方法
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// 请求消息
message HelloRequest {
    string name = 1;
}

// 响应消息
message HelloResponse {
    string reply = 1;
}

~~~

## 编写Server端GO代码 

1. 新建一个hello_server项目，根目录下执行go mod init hello_server
2. 新建一个`pb`文件夹，将上面的proto文件保存为hello.proto，将`go_pakcage` 填写Go代码的导入路径 

~~~go
// ...

option go_package = "hello_server/pb";

// ...

~~~

此时项目路径 

~~~go
hello_server
├── go.mod
├── go.sum
├── main.go
└── pb
    └── hello.proto

~~~

生成go代码 ：

~~~go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./hello_serve/pb/hello.proto 
~~~

生成的go源码文件会保存在pb文件夹下：

~~~go
hello_server
├── go.mod
├── go.sum
├── main.go
└── pb
    ├── hello.pb.go
    ├── hello.proto
    └── hello_grpc.pb.go

~~~

### 编写业务逻辑 

~~~go
package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"

	"google.golang.org/grpc"
)


type server struct{
	pb.UnimplementedGreeterServer
}


func (s *server) SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloResponse,error){
	return &pb.HelloResponse{Reply: "hello"+in.Name,},nil
}




func main(){
	lis, err := net.Listen("tcp",":8972")
	if err != nil{
		fmt.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()  //创建grpc服务器 
	pb.RegisterGreeterServer(s,&server{}) //在grpc服务端注册服务 

	//启动服务 
	err = s.Serve(lis)
	if err != nil{
		fmt.Printf("failed to server :%v",err)
		return 
	}
}
~~~

### 解析

**实现gRPC服务端接口 **

~~~go
type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloResponse,error){
	return &pb.HelloResponse{Reply: "hello"+in.Name,},nil
}
~~~

* 定义了一个`server`结构体，通过匿名字段继承了`pb.UnimplementedGreeterServer`,以便于注册服务时实现全部的`grpc`服务端接口
* 实现了`Sayhello`方法，实现了`_grpc.pb.go`的文件中的服务端的接口 

**创建gRPC服务器并注册服务**

~~~go
lis, err := net.Listen("tcp",":8972")
if err != nil{
	fmt.Printf("failed to listen: %v", err)
}

s := grpc.NewServer()  //创建grpc服务器 
pb.RegisterGreeterServer(s,&server{}) //在grpc服务端注册服务

~~~

* 首先使用 `net.Listen` 函数创建一个 TCP 监听器
* 通过 `grpc.NewServer()` 创建一个 gRPC 服务器。
* 最后，调用 `pb.RegisterGreeterServer(s, &server{})` 方法来将 `server` 结构体中实现的所有 gRPC 服务端接口注册到 gRPC 服务器上。

**启动gRPC服务端提供服务**

~~~go
err = s.Serve(lis)
if err != nil{
	fmt.Printf("failed to server :%v",err)
	return 
}
~~~

* 启动gRPC服务器，并提供服务 



## 编写Client客户端

1. 新建一个`hello_client`项目，在项目根目录下执行 go mod init hello_client

2. 新建pb文件夹，将上面proto文件保存为hello.proto，将`go_package`按如下方式修改。

   ```protobuf
   // ...
   
   option go_package = "hello_client/pb";
   
   // ...
   ```

生成go代码 ：

~~~go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto

~~~

此时项目目录：

~~~go
http_client
├── go.mod
├── go.sum
├── main.go
└── pb
    ├── hello.pb.go
    ├── hello.proto
    └── hello_grpc.pb.go

~~~

### 编写业务逻辑

~~~go
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"hello_client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
~~~

### 解析 

**声明命令行参数，并解析命令行参数**

~~~go
const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
}
~~~

* 声明了两个命令行参数：`addr`和`name`,分别表示服务端地址和欢迎词中的名字，通过`flag.Parse()`方法来解析命令行参数

**连接到gRPC服务器并创建客户端 **

~~~go
conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
if err != nil {
    log.Fatalf("did not connect: %v", err)
}
defer conn.Close()
c := pb.NewGreeterClient(conn)

~~~

* `grpc.Dial`方法连接到gRPC服务器。并使`insecure.NewCredentials()` 方法来创建不安全的 TLS 凭证。
* 通过 `pb.NewGreeterClient(conn)` 创建一个 gRPC 客户端。

**调用gRPC服务端方法并打印响应数据 **

~~~go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
if err != nil {
    log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetReply())

~~~

* 使用 `context.WithTimeout` 方法创建一个带有超时时间的上下文对象，并在结束时调用 `cancel` 函数以释放相关资源。
* 调用 `c.SayHello` 方法来向 gRPC 服务端发送请求，并将响应结果赋值给 `r` 变量。
  * 如果请求失败，则输出错误信息并退出程序。
  * 如果请求成功，则通过 `r.GetReply()` 方法获取服务端返回的欢迎词并将其打印出来。





# grpc流式开发 

上面例子中客户端发送一个RPC请求到服务端，服务端进行业务处理并返回响应数据给客户端，这是gRPC最基本的一种工作方式。

依托于`HTTP2`，`gRPC`还支持流式`RPC`

## 服务端流式RPC

客户端发出一个RPC请求，服务端与客户端之间建立一个单向的流

服务端可以向流中写入多个响应信息，最后主动关闭流，而客户端需要监听这个流，不断获取响应直到流关闭

**定义服务**

~~~go
// 服务端返回流式数据
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
~~~

修改`.proto`文件后，需要重新使用 protocol buffers编译器生成客户端和服务端代码。

**服务端实现`LotsOfReplies方法 `**

~~~go
// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

~~~

**客户端调用`lotsOfReplies`并将收到的数据依次打印出来**

~~~go
// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

~~~

## 客户端流式RPC

客户端流传入多个请求对象，服务端返回一个响应结果，

典型案例：物联网终端向服务器上报数据，大数据流式计算 

在这个示例中:编写一个多次发送人名，服务端统一返回一个打招呼的程序 

**定义服务 **

~~~go
// 客户端发送流式数据
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
~~~

**服务端实现LotOfGreetings方法**

~~~go
// LotsOfGreetings 接收流式数据
func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好："
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += res.GetName()
	}
}  

~~~

**客户端调用`LotsOfGreetings`方法**

向服务端发送流式请求数据，接收返回值后打印 

~~~go
func runLotsOfGreeting(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式RPC
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
	}
	names := []string{"七米", "q1mi", "沙河娜扎"}
	for _, name := range names {
		// 发送流式数据
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err: %v", name, err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed: %v", err)
	}
	log.Printf("got reply: %v", res.GetReply())
}
~~~

## 双向流式RPC

双向流式RPC即客户端和服务端均为流式的RPC，能发送多个请求对象也能接收到多个响应对象。典型应用示例：聊天应用等。

我们这里还是编写一个客户端和服务端进行人机对话的双向流式RPC示例。

**定义服务 **

~~~go
// 双向流式数据
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);

~~~

**服务端实现`BindiHello`****方法** 

~~~go
// BidiHello 双向流式打招呼
func (s *server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := magic(in.GetName()) // 对收到的数据做些处理

		// 返回流式响应
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}

~~~

### 客户端

**客户端调用BindiHello方法**

一边从总段获取输入的请求数据发送到服务算，一边从服务端接收流式响应 

~~~go
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
	<-waitc
}

~~~

#### 客户端解析

**创建上下文对象和双向流式连接**

~~~go
ctx,cancle := context.WithTimeout(context.Background(),2*time.Second)
defer cancle()

stream,err := c.BidiHello(ctx)
if err != nil{
    log.Fatalf("c.BidiHello failed ,err: %v",err)
}
~~~

* 首先使用`context.WithTimeout` 函数创建一个带有超时时间的上下文对象，并在结束时调用 `cancel` 函数以释放资源。
* 然后，我们使用 `c.BidiHello` 方法创建一个双向流式连接。

**执行流程 **

1. 主程序创建上下文对象和 双向流式连接，并启动一个goroutine来接收服务端返回的响应 
2. 主函数通过读取标准输入获取用户输入，并使用双向流式连接将数据发送到服务端。如果用户输入 `QUIT`，则主函数会关闭双向流式连接并等待 goroutine 执行完毕。
3. 在 goroutine 中，我们通过 `stream.Recv()` 方法不断地从服务端接收数据，并通过 `in.GetReply()` 方法获取数据内容。如果读取完毕，则关闭 `waitc` 通道退出 goroutine。
4. 当主函数接收到用户输入 `QUIT` 后，它会通过 `stream.CloseSend()` 关闭双向流式连接，然后等待 goroutine 执行完毕。在等待过程中，主函数会通过 `<-waitc` 语句从 `waitc` 通道中读取数据，以保证 goroutine 已经完成任务。
5. 整个流程顺序：主函数创建连接并发送数据到服务端--->goroutine接收服务端返回的响应--->主函数关闭连接并等待goroutine执行完毕，

无缓冲通道是一种阻塞式的通道，它的容量为 0。

* 当向通道发送数据时，发送操作会被阻塞，直到有其他 goroutine 从该通道中接收数据；
* 当从通道接收数据时，接收操作也会被阻塞，直到有其他 goroutine 向该通道中发送数据。
* 因此，在使用无缓冲通道时，我们不需要显式地使用锁来控制并发访问，通道的阻塞和唤醒机制就可以自动地保证并发安全。
