package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

func say(wg *sync.WaitGroup,page int)  {
	url := "https://so.gushiwen.org/shiwen/default_2Ab90660e3e492A"+strconv.Itoa(page)+".aspx"
	resp,_ := http.Get(url)
	data,_ := ioutil.ReadAll(resp.Body)
	html := string(data)
	exp := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
	results := exp.FindAllStringSubmatch(html,-1)
	var title string
	for _,result := range results{
		title = result[1]
	}
	fmt.Println("第"+strconv.Itoa(page)+"页:"+title)
	//fmt.Println(result)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i:=0;i<=100;i++{
		wg.Add(1)
		go say(&wg,i)
	}
	wg.Wait()
	fmt.Println("over")
}
