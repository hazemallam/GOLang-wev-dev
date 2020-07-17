package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("main.gohtml"))

}

func main() {
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
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	fileCreated, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer fileCreated.Close()

	err = tpl.ExecuteTemplate(fileCreated, "main.gohtml", data)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
