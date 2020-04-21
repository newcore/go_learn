package main

import "fmt"

func main() {
	m := make(map[int]int ,10)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m[4])
	m2 := map[int]int{1:3,2:9}
	fmt.Println(m2)
	fmt.Println(m2[99])
	v,ok := m2[4]
	fmt.Println(v,ok)

	m3 := map[string]string{}
	m3["name"] = "wf"
	v3,ok := m3["name1"]
	fmt.Println(v3,ok)
	fmt.Println("name2:"+m3["name2"])
}
