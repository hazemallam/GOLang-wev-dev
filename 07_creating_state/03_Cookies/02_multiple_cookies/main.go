package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/multiple", multiple)
	http.HandleFunc("/read", read)
	http.HandleFunc("/readall", readAll)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":5050", nil))

}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "cookie-1",
		Value: "cookie number one",
	})
	fmt.Fprintln(res, "COOKIE WRITEN-CHECK YOUR BROWSER")
}

func multiple(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "cookie-2",
		Value: "cookie number two",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "cookie-3",
		Value: "cookie number three",
	})
	fmt.Fprintln(res, "COOKIES WRITEN-CHECK YOUR BROWSER")
}

func read(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie-1")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #1 : ", cookie)
	}

	cookie, err = req.Cookie("cookie-2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #2 : ", cookie)
	}

	cookie, err = req.Cookie("cookie-3")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #3 : ", cookie)
	}
}

func readAll(res http.ResponseWriter, req *http.Request) {
	c := req.Cookies()
	for i, v := range c {
		fmt.Fprintf(res, "YOUR COOKIE %v# = %v\n\n", i, v)
	}
}
