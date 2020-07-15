package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/gold.jpg", dogPic)
	http.ListenAndServe(":5050", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/gold.jpg">`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "gold.jpg")
}
