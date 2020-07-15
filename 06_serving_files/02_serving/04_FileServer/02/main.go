package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":5050", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/resources/gold.jpg>`)
}
