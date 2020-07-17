package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"tf": TimeFormating,
}

func TimeFormating(t time.Time) string {
	return t.Format("02-01-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("main.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", time.Now())
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
