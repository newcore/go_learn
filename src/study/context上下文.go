package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("好了，子协程结束了")
				return
			default:
				fmt.Println("监控中")
				time.Sleep(1*time.Second)
			}
		}
	}()
	fmt.Println("开始干活，暂停5秒")
	time.Sleep(5*time.Second)
	cancel()
	fmt.Println("让子协程结束")
	time.Sleep(5*time.Second)
	fmt.Println("over")

}
