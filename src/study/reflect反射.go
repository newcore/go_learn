package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []string
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(),typeOfA.Kind())
	if nil == a{
		fmt.Println("a 是空的")
	}
}