package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}
func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/images/", fs)
	http.HandleFunc("/", dog)
	log.Fatal(http.ListenAndServe(":5050", nil))

}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error occured during executing file", err)
	}

}
