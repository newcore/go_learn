package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func watch(ctx context.Context,name string){
	for{
		select {
		case <-ctx.Done():
			fmt.Println(name+"结束了")
			return
		default:
			fmt.Println(name+"正在进行中")
			time.Sleep(1*time.Second)
		}
	}
}
func main() {
	ctx,cancel := context.WithCancel(context.Background())
	go watch(ctx,"监控1")
	go watch(ctx,"监控2")
	go watch(ctx,"监控3")
	time.Sleep(5*time.Second)
	cancel()
	fmt.Println("通知子协程结束")
	time.Sleep(3*time.Second)
	fmt.Println("over")
	log.Println("测试")
}
