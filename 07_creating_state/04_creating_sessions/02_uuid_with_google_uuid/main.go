package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
			return
		}
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
	}
	http.SetCookie(res, cookie)
	fmt.Println(cookie)
}
