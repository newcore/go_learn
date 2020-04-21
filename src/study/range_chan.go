package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,6)
	go func() {
		//time.Sleep(time.Second)
		//ch <- 1
		//time.Sleep(5*time.Second)
		//ch <- 2
		//fmt.Println("2发送完毕")
		//time.Sleep(3*time.Second)
		//close(ch)
		for i:=0;i<=10;i++{
			ch<-i
		}
		time.Sleep(2*time.Second)
		close(ch)
	}()

	for c := range ch{
		fmt.Println(c)
	}
}
