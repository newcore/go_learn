package main

import (
	"errors"
	"fmt"
)

func div(num int) (int,error)  {
	if num == 0{
		return 0,errors.New("除数不能为0")
	}
	return 10/num,nil
}
func main() {
	result ,err := div(8)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result,err = div(0)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
