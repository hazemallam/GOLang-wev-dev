package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":5050", nil)

}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Fatalln("error occured", err)
	}

}
