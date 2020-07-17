package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("error in parseGlob", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("error in Execute", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln("error in ExecuteTemplate", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln("error in ExecuteTemplate", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln("error in ExecuteTemplate", err)
	}
}
