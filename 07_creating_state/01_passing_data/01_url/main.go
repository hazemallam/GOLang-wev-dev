package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	value := req.FormValue("q")
	io.WriteString(res, "search value "+value)
}
