package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
	fmt.Println("helloword")
}

func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8081", nil)
}