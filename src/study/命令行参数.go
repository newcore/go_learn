package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for i,arg := range args{
		fmt.Printf("第%d个参数的值是:%s",i,arg)
	}
}
