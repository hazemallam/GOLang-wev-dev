package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (h hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog", d)
	http.Handle("/cat", c)
	http.ListenAndServe(":5050", nil)
}
