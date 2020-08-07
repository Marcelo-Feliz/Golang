package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name  string
	Meal  string
	Price float64
}

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := items{
		item{
			Name:  "Torrada",
			Meal:  "Breakfast",
			Price: 4.95,
		},
		item{
			Name:  "Hamburger",
			Meal:  "Lunch",
			Price: 6.95,
		},
		item{
			Name:  "Pizza",
			Meal:  "Dinner",
			Price: 7.95,
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
