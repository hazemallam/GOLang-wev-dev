package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply", apply)
	log.Fatalln(http.ListenAndServe(":5050", nil))
}

func index(res http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	HandleError(res, err)
}

func about(res http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.html", nil)
	HandleError(res, err)
}

func contact(res http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(res, "contact.html", nil)
	HandleError(res, err)
}

func apply(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(res, "applyProcess.html", nil)
		HandleError(res, err)
	}
	err := tpl.ExecuteTemplate(res, "apply.html", nil)
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
