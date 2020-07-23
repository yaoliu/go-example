//__author__ = "YaoYao"
//Date: 2019-06-16
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

type UserRequest struct {
	Id int `json:"id"`
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct{}

func (u *User) GetUserInfo(req *UserRequest, resp *UserResponse) error {
	resp.Id = 1
	resp.Name = "Yaoyao"
	resp.Age = 18
	return nil
}

func main() {
	_ = rpc.RegisterName("Rect", new(Rect))
	_ = rpc.RegisterName("User", new(User))
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
		go jsonrpc.ServeConn(conn)
	}
}
