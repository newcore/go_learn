package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func getPage(page int,ch chan int){
	fmt.Printf("当前采集第%d页\n",page)
	url := "http://www.qncye.com/qibu/jihui/list_99_"+strconv.Itoa(page)+".html"
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println("err",err)
	}
	defer resp.Body.Close()
	html := ""
	buffer := make([]byte,1024*4)
	for {
		n,_ := resp.Body.Read(buffer)
		if n == 0 {
			//fmt.Println(err1)
			break
		}
		html += string(buffer[:n])
	}
	html = ConvertToString(html,"gbk","utf-8")
	exp := regexp.MustCompile(`target=\"_blank\" title=\"(.*?)\">.*?</a></h3>`)
	contentExp := regexp.MustCompile(`<p class="txt">(.*?)</p>`)
	titleResult := exp.FindAllStringSubmatch(html,-1)
	contentResult := contentExp.FindAllStringSubmatch(html,-1)
	count := len(titleResult)
	fmt.Println(len(titleResult))
	fmt.Println(len(contentResult))
	for i:=0;i<count;i++{
		title := titleResult[i][1]
		content := contentResult[i][1]
		fmt.Println("标题:",title)
		fmt.Println("内容:",content)
		fmt.Println(strings.Repeat("---",30))
	}

	//fmt.Println(html)
	go save_txt(page,html)
	ch<-page
}

func save_txt(page int,content string){
	os.Mkdir("page",os.ModePerm)
	root,_ := os.Getwd()
	filename := root + "/page/" + strconv.Itoa(page) + ".txt"
	f,err := os.Create(filename)
	if err != nil{
		fmt.Println("创建文件失败",err)
		return
	}
	defer f.Close()
	f.WriteString(content)
	fmt.Println("写入文件成功")
}
func main() {
	start := time.Now()
	var ch = make(chan int)
	for i:=1;i<=100;i++{
		go getPage(i,ch)
	}
	for i:=1;i<=100;i++{
		<-ch
	}
	fmt.Println(time.Now())
	fmt.Println("一共耗时",time.Now().Sub(start))
	fmt.Println("搞定")
}
