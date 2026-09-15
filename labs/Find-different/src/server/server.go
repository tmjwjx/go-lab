package main

import (
	"Find-different/models"
	"fmt"
	"html/template"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/house.html")
	t.Execute(w, models.User{Id: 12})
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Gopher!")
}

func main() {

	server := http.Server{Addr: "localhost:8899"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	http.HandleFunc("/sayHello", sayHello) // 设置访问的路由
	server.ListenAndServe()
}