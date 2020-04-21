package main

import "fmt"

func main() {
	ch := make(chan int,3)
	done := make(chan bool)

	go func() {
		for{
			msg,ok := <-ch
			if ok{
				fmt.Println("接收到",msg)
			} else {
				fmt.Println("没有值了")
				done<-true
				return
			}
		}
	}()
	for i:=0;i<3;i++{
		ch<-i
		fmt.Println("发送了",i)
	}
	close(ch)
	<-done
}
