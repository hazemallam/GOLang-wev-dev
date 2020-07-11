package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotes  []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {

	r := region{
		Region: "Southern",
		Hotes: []hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
