package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)

}

func set(res http.ResponseWriter, req *http.Request){
	http.SetCookie(res, &http.Cookie{
		Name: "my-cookie",
		Value: "some-value",
	})
	fmt.Fprintln(res, "COOKIE WRITEN-CHECK YOUR BROWSER")
}

func read(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNoContent)
		return
	}
	fmt.Fprintln(res, "YOUR COOKIE: ", cookie)
}