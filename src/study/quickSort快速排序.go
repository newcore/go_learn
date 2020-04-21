package main

import "fmt"

func quickSort(scores []int)  []int{
	for i:=0;i<len(scores)-1;i++{
		fmt.Printf("第%d轮\n",i)
		for j:=i+1;j<len(scores);j++{
			if scores[i] > scores[j]{
				scores[i],scores[j] = scores[j],scores[i]
				fmt.Println("++++++",scores)
			}
		}
		fmt.Println(scores)
	}
	return scores
}
func main() {
	scores := []int{1,9,8,50,40,30,99,88,4,56}
	fmt.Println("初始值",scores)
	scores = quickSort(scores)
	fmt.Println(scores)
}