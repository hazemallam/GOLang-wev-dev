package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {

	xs := []string{"Zero", "one", "two", "three", "four"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
