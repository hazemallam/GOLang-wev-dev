package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveFile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)
}

func serveFile(res http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		//open file
		file, header, err := req.FormFile("cv")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Println("\nfile:", file, "\nheader:", header, "\nerror", err)

		//read file
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	res.Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="cv">
			<input type="submit">
		</form>

	`+s)
}
