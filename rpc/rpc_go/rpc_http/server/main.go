//__author__ = "YaoYao"
//Date: 2019-06-16
package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err == nil {
		listener, err := net.Listen("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("ing")
			go rpc.ServeConn(conn)
		}
	}
}
