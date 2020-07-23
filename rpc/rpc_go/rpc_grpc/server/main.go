//__author__ = "YaoYao"
//Date: 2019-06-17
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc_go/rpc_grpc/pb"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "x",}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err == nil {
		log.Fatal(err)
	}
}
