package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/gold.jpg", imge)
	log.Fatal(http.ListenAndServe(":5050", nil))

}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("main.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "main.gohtml", nil)
}

func imge(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./images/gold.jpg")
}
