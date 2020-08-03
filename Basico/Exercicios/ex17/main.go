package main

import "fmt"

func main() {

	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))         //Print “s” converted to a slice of byte.
	fmt.Println(string([]byte(s))) //Print “s” converted to a slice of byte and then converted back to a string.
	fmt.Println(s[:14])            //Print dos primeiros 14 carateres
	fmt.Println(s[10:22])
	fmt.Println(s[17:])
	for _, i := range s {
		fmt.Println(string(i))
	}
}
