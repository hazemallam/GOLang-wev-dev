package main

import (
	"html/template"
	"os"
	"log"
)

type page struct{
	Title, Heading, Input string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main(){
	home := page{
		Title: "Escaped",
		Heading: "Danger is escaped with html/template",
		Input: `<script>alert("you");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", home)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}