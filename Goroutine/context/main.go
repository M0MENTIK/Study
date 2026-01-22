package main

import (
	"context"
	"fmt"
	"time"
)

// parents context
func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo finish")
			return
		default:
			fmt.Println("foo", n)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// dauther context
func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo finish")
			return
		default:
			fmt.Println("boo", n)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)
	go boo(childContext, 1)
	go boo(childContext, 2)
	go boo(childContext, 3)

	time.Sleep(1 * time.Second)
	childCancel() // отмена childcontext

	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(3 * time.Second)
}
