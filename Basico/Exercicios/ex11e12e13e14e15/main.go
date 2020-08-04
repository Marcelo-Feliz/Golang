package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, i am gator")
}
func main() {
	var g1 gator
	g1 = 5
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	//x = g1
	// you CANNOT assign g1 to x. They are two different types!

	x = int(g1)
	fmt.Println(x)
	g1.greeting()
}
