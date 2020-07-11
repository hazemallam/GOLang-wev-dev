package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  double,
	"fsqr":  square,
	"fsqrt": sqroot,
}

func double(n int) float64 {
	return float64(n) * 2
}

func square(n float64) float64 {
	return math.Pow(float64(n), 2)
}

func sqroot(n float64) string {
	s := fmt.Sprintf("%.2f", math.Sqrt(float64(n)))
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("main.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", 3)
	if err != nil {
		log.Fatalln("error occured", err)
	}
}
