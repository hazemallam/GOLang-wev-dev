package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
