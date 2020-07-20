package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)
}

func foo(res http.ResponseWriter, req *http.Request){
	fmt.Print("your request method at foo: ",req.Method,"\n\n")
}

func bar(res http.ResponseWriter, req *http.Request){
	fmt.Print("your request method at bar: ", req.Method, "\n\n")
	//process from submission
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func barred(res http.ResponseWriter, req *http.Request){
	fmt.Print("your request method at barred: ", req.Method, "\n\n")
	tpl.ExecuteTemplate(res, "index.html", nil)
}