package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip, Meal string
	Price               float64
}

var tpl *template.Template

type items []item

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	i := items{
		item{
			Name:    "Oatmeal",
			Descrip: "yum yum",
			Meal:    "Breakfast",
			Price:   4.95,
		},
		item{
			Name:    "Hamburger",
			Descrip: "Delicous good eating for you",
			Meal:    "Lunch",
			Price:   6.95,
		},
		item{
			Name:    "Pasta Bolognese",
			Descrip: "From Italy delicious eating",
			Meal:    "Dinner",
			Price:   7.95,
		},
	}
	err := tpl.Execute(os.Stdout, i)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
