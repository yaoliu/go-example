//__author__ = "YaoYao"
//Date: 2020/6/29
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	go func() {
		for {
			ch1 <- 1
			ch2 <- 2
		}
	}()
	select {
	case <-ch1:
		fmt.Println(1)
	case <-ch2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}
