package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {

	p := person{
		Name: "James",
		Age:  42,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
