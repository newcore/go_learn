package main

import (
	"fmt"
	"time"
)

func print(ch chan int,quit chan int)  {
	for{
		select {
		case num :=<- ch:
			fmt.Println("接收到了",num)
		case <-time.After(time.Second*10):
			fmt.Println("超时了")
			quit<-1
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan int)
	go print(ch,quit)
	for i:=0;i<=5;i++{
		fmt.Println("发送",i)
		ch<-i
		time.Sleep(time.Second)
	}
	//<-quit
	fmt.Println("结束")
}
