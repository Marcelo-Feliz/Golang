package main

import "fmt"

func main() {
	i := 7
	inc(&i)
	fmt.Println(&i) //print do endereço do pointer
	fmt.Println(i)  //print do valor da variável
}

func inc(x *int) {
	*x++ //pointer
}
