//__author__ = "YaoYao"
//Date: 2019-06-16
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type UserRequest struct {
	Id int `json:"id"`
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	var ret int

	err = rpc.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	var userResp UserResponse
	err = rpc.Call("User.GetUserInfo", UserRequest{1}, &userResp)
	fmt.Println(userResp.Name)
}
