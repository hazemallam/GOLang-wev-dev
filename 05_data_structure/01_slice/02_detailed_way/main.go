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

	s := []string{"James", "Bond", "Miss", "Moneypenny"}
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
