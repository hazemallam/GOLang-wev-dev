package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("gold.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	http.ServeContent(res, req, fi.Name(), fi.ModTime(), f)
}
