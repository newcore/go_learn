package main

import (
	"os"
)

func main()  {
	os.Mkdir("page1",os.ModePerm)
	//path,_ := os.Getwd()
	//fmt.Println(path)
}
