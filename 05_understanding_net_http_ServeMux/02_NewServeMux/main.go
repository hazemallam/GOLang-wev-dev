package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dogy dogy dogy")
}

type hotcat int

func (h hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // go to any thing after dog/ in url
	mux.Handle("/cat", c)

	http.ListenAndServe(":5050", mux)
}
