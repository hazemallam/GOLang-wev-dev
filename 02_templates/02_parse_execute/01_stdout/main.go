package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	template, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln("error occured", err)
	}

	err = template.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("error occured", err)
	}

}
