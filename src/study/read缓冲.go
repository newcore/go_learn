package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("hello golang, i love it now!")
	b := make([]byte,8)
	result := ""
	for{
		n,err := r.Read(b)
		if n == 0{
			break
		}
		fmt.Println(n,string(b[:n]),err,string(b))
		result += string(b[:n])
	}
	fmt.Println(result)
}
