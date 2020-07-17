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

type doubleZero struct {
	person
	LicenceToKill bool
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	p1 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  42,
		},
		LicenceToKill: false,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
