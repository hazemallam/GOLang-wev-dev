package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", serveFile)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)
}

func serveFile(res http.ResponseWriter, req *http.Request) {
	// var s string
	if req.Method == http.MethodPost {
		//open
		file, header, err := req.FormFile("cv")
		if err != nil {
			http.Error(res, err.Error(), 5000)
			return
		}
		defer file.Close()
		fmt.Println("\nfile", file, "\nheader", header, "\nerror", err)

		//read
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		// s = string(bs)

		//store on server
		dst, err := os.Create(filepath.Join("./user/", header.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = dst.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	res.Header().Set("content-type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.html", nil)

}
