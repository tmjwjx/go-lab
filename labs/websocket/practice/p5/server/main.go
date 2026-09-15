package main

import "net/http"

type MyHander struct {
}
type MyHandle struct {
}

func (m *MyHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("返回的数据哈哈"))
}
func (m *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MyHandle返回的数据哈哈"))
}
func main() {
	//单控制器
	h := MyHander{}
	h2 := MyHandle{}
	server := http.Server{Addr: "localhost:8090"}
	http.Handle("/first", &h)
	http.Handle("/second", &h2)
	server.ListenAndServe()
}