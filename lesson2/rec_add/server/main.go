package main

import (
	"log"
	"net"
	"net/rpc"
)

//基于rpc实现的服务端

type Args struct {
	X, Y int
}

// SericeA自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(arg *Args, reply *int) error {
	*reply = arg.X + arg.Y
	return nil
}

// 以下代码定义ServiceA类型注册为一个服务，其Add方法就支持RPC调用
func main() {
	service := new(ServiceA)

	rpc.Register(service) //注册RPC服务
	//rpc.HandleHTTP()       //基于http协议

	l, e := net.Listen("tcp", ":9091") //设置监听地址
	if e != nil {
		log.Fatal("listen err:", e)
	}
	//http.Serve(l,nil)  //启动服务并监听请求---基于http

	//基于tcp  去除掉http层后，直接使用rpc.Serverconn进行调用
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
