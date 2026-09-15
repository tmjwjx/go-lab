package main

import (
	"fmt"
	"net/http"
)

func first(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "多函数-first")
}

func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "多函数-second")
}

func main() {
	//多控制器
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
	server.ListenAndServe()

}