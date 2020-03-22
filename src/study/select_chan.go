package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(5 * time.Second)
		c2 <- 2
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("get value :%d from c1\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("get value :%d from c2\n", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("3 second ...")
		default:
			fmt.Println("default value")
		}
	}
}
