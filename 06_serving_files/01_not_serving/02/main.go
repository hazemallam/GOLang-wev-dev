package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	http.ListenAndServe(":5050", nil)

}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<!--image doesn't serve-->
	<img src="/gold.jpg">
	`)
}
