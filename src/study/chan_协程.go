package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int,10)
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C{
			select {
			case <-done:
				fmt.Println("done !!!")
				return
			default:
				fmt.Printf("send message %d\n",<-message)
			}
		}
	}()
	for i:=0;i<10;i++{
		message<-i
	}
	time.Sleep(4*time.Second)
	close(done)
	time.Sleep(time.Second)
}
