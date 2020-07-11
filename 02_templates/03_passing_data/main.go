package main

import (
	"os"
	"log"
	"text/template"
	
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", 42)
	if err != nil{
		log.Fatalln("error", err)
	}
}