package main

import (
	"fmt"
	"net/http"
)

func param(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, h["Accept-Language"])

	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Form["name"])

	fmt.Fprintln(w, r.FormValue("age"))

}

func main() {
	//获取请求头和请求参数
	//localhost:8899/param?name=zhang&age=18
	server := http.Server{Addr: ":8899"}
	http.HandleFunc("/param", param)
	server.ListenAndServe()
}