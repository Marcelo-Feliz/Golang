package main

import "fmt"

func main() {

	m := map[string]int{
		"Todd":    45,
		"Job":     42,
		"Andre":   32,
		"Augusto": 18,
	}
	fmt.Println(m)

	for i := range m {
		fmt.Println(i)
	}
	for i, v := range m {
		fmt.Println(i, v)
	}

}
