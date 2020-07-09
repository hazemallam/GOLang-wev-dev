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

	xs := []string{"zero", "one", "two", "three", "four"}

	data := struct {
		Words []string
		Name  string
	}{
		xs,
		"James",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
