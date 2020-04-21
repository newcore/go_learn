package main

import "fmt"

type Person struct {
	name string
	age int
}
func main() {
	p := Person{name:"lisi",age:18}
	p.name = "王五"
	fmt.Println(p)
	param := &p
	(*param).age = 30
	fmt.Println(*param)
}
