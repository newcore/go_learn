package main

import "fmt"

type Phone interface {
	call()
}
type Xiaomi struct {

}

func (xiaomi Xiaomi) call() {
	fmt.Println("it is xiaomi ,it can call")
}

type Iphone struct {

}

func (iphone Iphone) call() {
	fmt.Println("this is Iphone")
}
func main() {
	xiaomi := new(Xiaomi)
	xiaomi.call()
	iphone := new(Iphone)
	iphone.call()
}
