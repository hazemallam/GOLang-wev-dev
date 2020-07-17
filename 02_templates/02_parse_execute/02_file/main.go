package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templatee, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln("error occured", err)
	}
	createdFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error occured", err)
	}
	defer createdFile.Close()
	err = templatee.Execute(createdFile, nil)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
