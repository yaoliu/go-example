//__author__ = "YaoYao"
//Date: 2020/6/23
package main

import "fmt"

func Default() {
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(2)
	}()

	defer func() {
		fmt.Println(3)
	}()
}

func main() {
	Default()
}
