package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user struct {
	FirstName string
	LastName  string
	Sub       bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", serveUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":5050", nil))
}

func serveUser(res http.ResponseWriter, req *http.Request) {
	fname := req.FormValue("fname")
	lname := req.FormValue("lname")
	subscribe := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(res, "index.html", user{fname, lname, subscribe})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
