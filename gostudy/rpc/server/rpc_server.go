package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// var code string = "tcp"
var code string = "http"

type Service struct{}

//满足go的rpc方法: 必须有2个可序列化的参数 第二个参数必须为指针
func (p *Service) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	if code == "tcp" {
		//Service 为一个domain ，domain中会包含可执行的函数
		rpc.RegisterName("Service", new(Service))

		listener, err := net.Listen("tcp", ":1234")

		if err != nil {
			log.Fatal("ListenTCP error", err)
		}
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error", err)
		}
		//处理请求
		rpc.ServeConn(conn)
	}
	if code == "http" {
		//如果不带domain 则以new的名字为准， 还是建议带
		// rpc.Register(new(Service))
		rpc.RegisterName("Service", new(Service))
		//http://127.0.0.1:1234/debug/rpc
		rpc.HandleHTTP()

		listener, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal("listen error", err)
		}
		http.Serve(listener, nil)
	}

}
