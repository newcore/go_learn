package main

import "fmt"

func intSeq() func() int{
	var i int
	return func() int {
		i++
		return i
	}
}
func main() {
	fmt.Println("第一次初始化intseq")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("第二次初始化intseq")
	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}
