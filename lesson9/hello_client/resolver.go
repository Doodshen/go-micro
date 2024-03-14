package main

import (
	"google.golang.org/grpc/resolver"
)

// 自定义name resolver

const (
	myScheme   = "kingshen"    
	myEndpoint = "resolver.kingshen.com"    
)
//通过调用myScheme把这个myEndpoint 转换为下面的addrs地址 

var addrs = []string{"127.0.0.1:8972", "127.0.0.1:8973", "127.0.0.1:8974"}

//自定义name resolver，实现Resolver接口
type kingshenResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *kingshenResolver) ResolveNow(o resolver.ResolveNowOptions) {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrList := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrList[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrList})
}

func (*kingshenResolver) Close() {}

// q1miResolverBuilder 需实现 Builder 接口
type kingshenResolverBuilder struct{}

func (*kingshenResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &kingshenResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			myEndpoint: addrs,
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
func (*kingshenResolverBuilder) Scheme() string { return myScheme }

func init() {
	// 注册 q1miResolverBuilder
	resolver.Register(&kingshenResolverBuilder{})
}
