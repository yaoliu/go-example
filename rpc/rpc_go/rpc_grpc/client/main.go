//__author__ = "YaoYao"
//Date: 2019-06-18
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"rpc_go/rpc_grpc/pb"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		log.Fatal(err)
	}()
	c := pb.NewGreeterClient(conn)
	name := "liuyao"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.SayHelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", r.Message)

}
