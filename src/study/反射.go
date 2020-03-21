package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID int
	name string
}
func main() {
	user := User{}
	t := reflect.TypeOf(user)
	fmt.Println(t)
}