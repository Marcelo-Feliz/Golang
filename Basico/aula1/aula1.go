package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"todd": 45,
		"job":  42,
	}

	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(sa1)

	sa1.speak() //james bond fala pelo atributo secretAgent
	//drill to inner type
	sa1.person.speak() //james bond fala pelo atributo pessoa

	saySomething(p1)
	saySomething(sa1)
}
