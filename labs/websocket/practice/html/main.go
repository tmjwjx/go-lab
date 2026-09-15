package main

import (
	"html/template"
	"net/http"
	"time"
)

type User struct {
	Name string
	Age  int
}

func MyTransfer(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("practice/html/view/index.html")
	t.Execute(w, User{"张三", 18})
}
func getTime(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("practice/html/view/time.html")
	time1 := time.Now()
	t.Execute(w, time1)
}
func html(w http.ResponseWriter, r *http.Request) {
	fm := template.FuncMap{"mt": MyTransfer} //建立一个fm映射，表示 mt->MyTransfer
	t := template.New("time.html").Funcs(fm) //模板t，名称time.html,可以调用fm
	t, _ = t.ParseFiles("practice/html/view/time.html")
	time1 := time.Now()
	t.Execute(w, time1)
}
func main() {
	server := http.Server{Addr: ":8090"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("practice/html/static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/gettime", getTime)
	http.HandleFunc("/html", html)
	server.ListenAndServe()
}