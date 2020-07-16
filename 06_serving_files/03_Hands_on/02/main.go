package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./resources")))
	http.HandleFunc("/dog/", dog)
	log.Fatal(http.ListenAndServe(":5050", nil))

}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./resources/html/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "./resources/html/index.html", nil)
}
