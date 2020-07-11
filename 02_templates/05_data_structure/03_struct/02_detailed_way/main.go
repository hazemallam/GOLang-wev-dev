package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Country string
	Name    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	p1 := person{
		Country: "USA",
		Name:    "James",
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
