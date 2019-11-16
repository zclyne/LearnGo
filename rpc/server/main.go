package main

import (
	rpcdemo "learngo.com/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 把所定义的除法service注册到rpc上
	rpc.Register(rpcdemo.DemoService{})

	// 创建一个服务器，开始监听1234端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		// 执行方法，这里开了一个goroutine来执行，是为了不阻塞Accept()接收新的rpc请求
		go jsonrpc.ServeConn(conn)
	}
}
