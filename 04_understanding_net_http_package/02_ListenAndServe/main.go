package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "type anything you want here")
}
func main() {
	var d hotdog
	http.ListenAndServe(":5050", d)
}
