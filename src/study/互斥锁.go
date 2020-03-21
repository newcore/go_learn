package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex
	for i:=0;i<2;i++{
		go func() {
			for i := 10000;i>0;i--{
				lock.Lock()
				count++
				lock.Unlock()
			}
			fmt.Println(count)
		}()
	}
	fmt.Scanf("\n")
}
