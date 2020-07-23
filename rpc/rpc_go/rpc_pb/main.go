//__author__ = "YaoYao"
//Date: 2019-06-16
package main

import (
	"Demo/rpc_pb/pb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	user1 := pb.User{
		Id:   *proto.Int32(1),
		Name: *proto.String("Mike"),
	}
	user2 := pb.User{
		Id:   2,
		Name: "John",
	}
	users := pb.MultiUser{
		Users: []*pb.User{&user1, &user2},
	}
	data, err := proto.Marshal(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users.Users[0].GetName())
	fmt.Println(data)

	var target pb.MultiUser
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(target.GetUsers()[0].GetName())
}
