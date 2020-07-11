package main

import (
	"text/template"
	"log"
	"os"
)

type sage struct{
	Name string
	Motto string
}

type car struct {
	Manufacture string
	Model string
	Doors int
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main(){
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacture: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacture: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sags := []sage{b,g,m}
	cars := []car{f,c}

	data := struct{
		Wisdom []sage
		Transport []car
	}{
		Wisdom : sags,
		Transport : cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil{
		log.Fatalln("error occured", err)
	}
}