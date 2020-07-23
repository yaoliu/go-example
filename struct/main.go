//__author__ = "YaoYao"
//Date: 2020/7/2
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	user := &Person{
		Name: "liuyao",
		Age:  18,
		Sex:  1,
	}
	fmt.Println(user)
	typeof := reflect.TypeOf(user)
	typeof.
 }
