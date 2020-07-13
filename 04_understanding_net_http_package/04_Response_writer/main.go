package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("My-key", "This is from me")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":5050", d)
}
