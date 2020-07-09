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

	s1 := struct {
		Score1 int
		Score2 int
	}{
		9,
		7,
	}

	err := tpl.Execute(os.Stdout, s1)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
