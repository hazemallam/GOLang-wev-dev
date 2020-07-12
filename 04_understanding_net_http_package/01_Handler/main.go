package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {
	var d hotdog
	http.ListenAndServe(":5050", d)
}
