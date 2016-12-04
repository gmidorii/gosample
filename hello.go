package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	fmt.Printf("hello world!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}