package main

import (
	"fmt"
	"time"
)

func work(w int,jobs chan int,result chan int)  {
	for job := range jobs{
		time.Sleep(time.Second)
		fmt.Printf("当前的是%d,jobs接收到的是:%d\n",w,job)
		result <- w*2
	}

}
func main() {
	jobs := make(chan int,100)
	result := make(chan int,100)
	for i:= 1;i<=3;i++{
		go work(i,jobs,result)
	}
	for i:=1;i<=9;i++{
		jobs<-i
	}
	close(jobs)
	for i:=0;i<9;i++{
		<-result
	}
}
