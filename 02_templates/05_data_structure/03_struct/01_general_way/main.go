package main

import (
	"text/template"
	"os"
	"log"
)

type person struct {
	Country string //FIELD NAME MUST START WITH CAPITAL LETTER
	Name string

}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main(){
	p1 := person{
		Country : "USA", 
		Name : "James",
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil{
		log.Fatalln("error occured", p1)
	}
}