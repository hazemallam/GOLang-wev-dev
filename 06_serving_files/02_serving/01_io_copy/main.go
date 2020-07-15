package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/gold.jpg", dogpic)
	http.ListenAndServe(":5050", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<img src="/gold.jpg">
	`)

}

func dogpic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("gold.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(res, f)
}
