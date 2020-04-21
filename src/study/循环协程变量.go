package main

import (
	"fmt"
	"time"
)

func main() {
	s := []string{"one","two","three"}
	for _,v := range s{
		go func() {
			time.Sleep(1*time.Second)
			fmt.Println(v)
		}()
	}
	time.Sleep(3*time.Second)
}
