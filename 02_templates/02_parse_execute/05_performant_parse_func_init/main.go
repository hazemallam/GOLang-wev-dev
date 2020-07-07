package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init(){

	tpl = template.Must(template.ParseGlob("templates/*"))

}
func main(){
	err := tpl.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln("error occured", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil{
		log.Fatalln("error occured", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil{
		log.Fatalln("error occured", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil{
		log.Fatalln("error occured", err)
	}
}