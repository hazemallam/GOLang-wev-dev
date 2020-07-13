package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "dogy dogy dogy")
	case "/cat":
		io.WriteString(res, "cat cat cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":5050", d)
}
