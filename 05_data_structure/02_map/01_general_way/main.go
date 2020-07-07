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

	m := map[string]string{
		"USA": "James",
		"UAE": "Khaled",
		"UK":  "Bond",
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
