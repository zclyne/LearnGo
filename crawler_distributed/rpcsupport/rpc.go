package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 包装rpc服务器
func ServeRpc(host string, service interface{}) error {
	// 把获得的service注册到rpc上
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	// 创建一个服务器，开始监听host端口
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
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

// 包装rpc客户端
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}