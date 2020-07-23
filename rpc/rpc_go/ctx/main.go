//__author__ = "YaoYao"
//Date: 2019-06-20
package main

import (
	"context"
	"fmt"
	"time"
)

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// ……
}

func r() {
	ctx := context.Background()
	process(ctx)
	ctx = context.WithValue(ctx, "traceId", "yaoyao-2019")
	process(ctx)
}

func process(ctx context.Context) {
	v, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no")
	}
}
