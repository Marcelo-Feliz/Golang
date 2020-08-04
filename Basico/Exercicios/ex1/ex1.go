package main

import "fmt"

func main() {

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	for i := range xi {
		fmt.Println(i)
	}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
