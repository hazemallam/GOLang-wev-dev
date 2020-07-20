package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user struct {
	First, Last string
	Sub         bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":5050", nil))

}

func serve(res http.ResponseWriter, req *http.Request) {
	//body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)
	err := tpl.ExecuteTemplate(res, "index.html", body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
