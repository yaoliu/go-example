//__author__ = "YaoYao"
//Date: 2019-06-16
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "liuyao", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
