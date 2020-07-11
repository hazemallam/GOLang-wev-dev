package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil{
		log.Fatalln("error occured", err)
	}
}