package main

import "fmt"

func sum(num ...int) int {
	var total int
	for i,v := range num{
		fmt.Printf("the %d params value:%d\n",i,v)
		total += v
	}
	return total
}
func main() {
	r1 := sum(2,3)
	fmt.Println(r1)
	r2 := sum(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(r2)
}
