//__author__ = "YaoYao"
//Date: 2019-06-19
package main

import (
	"fmt"
	"reflect"
)

type service struct {
	name string        // name of service
	rcvr reflect.Value // receiver of methods for the service
	typ  reflect.Type  // type of the receiver
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u user) Insert() string {
	return u.Name
}

func main() {
	u := &user{Id: 1, Name: "liuyao", Age: 18}
	DoFiledAndMethod(*u)
	//s := new(service)
	//rcvr := new(Func)
	//s.typ = reflect.TypeOf(rcvr)
	//s.rcvr = reflect.ValueOf(rcvr)
	//fmt.Println(s.rcvr)
	//var num int = 1
	//fmt.Println(reflect.TypeOf(num))
	//fmt.Println(reflect.ValueOf(num))
}

func DoFiledAndMethod(r interface{}) {
	getType := reflect.TypeOf(r)
	fmt.Println("get type  is", getType)
	getValue := reflect.ValueOf(r)
	fmt.Println("get value is", getValue)
	//获取type
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("Field: %s: %v = %v\n", field.Name, field.Type, value)
	}
	//获取method
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("Method: %s: %v\n", m.Name, m.Type)
	}
	v := reflect.Indirect(getValue)
	fmt.Println(v)
}
