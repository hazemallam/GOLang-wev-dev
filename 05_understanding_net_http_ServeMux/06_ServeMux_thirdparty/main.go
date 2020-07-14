package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", bolgWrite)
	err := http.ListenAndServe(":5050", router)
	if err != nil {
		log.Fatalln("error on listennig to the server", err)
	}

}

func user(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "READ CATEGORY, %s\n", ps.ByName("category"))
	fmt.Fprintf(res, "READ ARTICLE, %s\n", ps.ByName("article"))
}

func bolgWrite(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "WRITE CATEGORY, %s\n", ps.ByName("category"))
	fmt.Fprintf(res, "WRITE ARTICLE, %s\n", ps.ByName("article"))
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	HandleError(res, err)
}

func about(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	HandleError(res, err)
}

func contact(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "contact.gohtml", nil)
	HandleError(res, err)
}

func apply(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "apply.gohtml", nil)
	HandleError(res, err)
}

func applyProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "applyProcess.gohtml", nil)
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
