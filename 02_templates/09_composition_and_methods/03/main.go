package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number, Name, Unit string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring semester
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
