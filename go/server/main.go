package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	registerHomeFunction()
	registerBlogFunction()

	http.ListenAndServe(":3000", nil)

	fmt.Println("Server started on port 3000")
}

func registerHomeFunction() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "<h1>Hello World!</h1>")
	})
}

func registerBlogFunction() {
	http.HandleFunc("/blog", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "<h1>Hello Blog!</h1>")
	})
}
