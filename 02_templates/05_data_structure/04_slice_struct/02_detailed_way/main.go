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
	p2 := person{
		Country: "UAE",
		Name:    "Khaled",
	}
	p3 := person{
		Country: "UK",
		Name:    "Bond",
	}
	persons := []person{p1, p2, p3}

	err := tpl.Execute(os.Stdout, persons)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
