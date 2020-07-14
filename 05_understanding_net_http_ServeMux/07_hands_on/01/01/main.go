package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "index index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "dog dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "me me")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":5050", nil)
}
