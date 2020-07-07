package main

import "fmt"

type person struct {
	fname string
	lname string
}
type secretAgent struct {
	person
	licenceToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning Miss."`)
}
func (s secretAgent) speak() {
	fmt.Println(s.fname, `says, "shaken not stirred."`, s.licenceToKill)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

var y string = "hello"

const w = 30

const (
	age  = 20
	name = "James"
)

func main() {
	x := 7
	const z = 20
	xi := []int{1, 2, 3, 4, 5}
	mi := map[string]int{
		"James": 40,
		"Tod":   50,
	}
	p1 := person{
		fname: "James",
		lname: "Bond",
	}
	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
		},
		licenceToKill: true,
	}
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)
	fmt.Println(xi)
	fmt.Printf("%T\n", xi)
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	p1.speak()
	sa1.speak()
	sa1.person.speak()
	fmt.Println("---------------------------------------")
	saySomething(p1)
	saySomething(sa1)
}
