package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func showUser(w http.ResponseWriter, r *http.Request) {
	u := make([]User, 0)
	u = append(u, User{"张三", 18})
	u = append(u, User{"李四", 20})
	u = append(u, User{"王五", 22})
	b, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(b))
}

func download(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename") // 获取文件名
	f, err := ioutil.ReadFile("D:/桌面/" + filename)
	if err != nil {
		fmt.Fprintln(w, "文件下载失败", err)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(f)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("name")                                               // 获取文件名
	file, fileHeader, _ := r.FormFile("file")                                     // 获取文件
	b, _ := ioutil.ReadAll(file)                                                  // 读取文件内容
	fileName += fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):] // 获取文件后缀
	ioutil.WriteFile("D:/桌面/"+fileName, b, 0666)                                  // 写入文件
	t, _ := template.ParseFiles("html/view/success.html")                         // 显示成功后的模板
	t.Execute(w, nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/view/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "mykey", Value: "myvalue"}
	http.SetCookie(w, &cookie)
	t, _ := template.ParseFiles("html/view/index.html")
	t.Execute(w, nil)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cs := r.Cookies()
	t, _ := template.ParseFiles("html/view/index.html")
	t.Execute(w, cs)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("html/static"))))
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/", welcome)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()
}
