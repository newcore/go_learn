package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello ,this is index page")
}
func main() {
	http.HandleFunc("/",index)
	http.ListenAndServe("127.0.0.1:8080",nil)
}
