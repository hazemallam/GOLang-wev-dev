package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "me me")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog")
}

func sendData(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("main.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "main.gohtml", "hello")
}

func main() {
	http.Handle("/", http.HandlerFunc(me))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/send", http.HandlerFunc(sendData))
	http.ListenAndServe(":5050", nil)
}
