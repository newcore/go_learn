package main

import "fmt"

type Stu struct {
	name string
	age int
	score int
}

func GrowUp(s *Stu)  {
	s.age += 1
	s.score += 10
}
func main() {
	s := new(Stu)
	s.age = 19
	s.score = 80
	s.name = "张三"
	fmt.Println(s)
	GrowUp(s)
	fmt.Println(s)

	var s2 Stu
	s2.name = "王五"
	fmt.Println(s2)
	GrowUp(&s2)
	fmt.Println(s2)
}
