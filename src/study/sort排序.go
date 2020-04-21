package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,5,4,6,7,3,55,2}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
