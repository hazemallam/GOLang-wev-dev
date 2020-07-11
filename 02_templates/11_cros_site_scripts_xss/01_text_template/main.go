package main

import (
	"log"
	"os"
	"text/template"
)

type page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	home := page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("You");</script>`,
	}
	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
