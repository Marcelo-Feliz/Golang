package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

type restaurante struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	data := []restaurante{
		restaurante{
			Name: "Fernando",
			Menu: menu{
				Breakfast: []string{
					"Torrada",
					"Leite",
				},
				Lunch: []string{
					"Bitoque",
					"Bacalhau com natas",
				},
				Dinner: []string{
					"Pizza",
					"Salada",
				},
			},
		},
		restaurante{
			Name: "A Grelha",
			Menu: menu{
				Breakfast: []string{
					"Torrada",
					"Leite",
				},
				Lunch: []string{
					"Bitoque",
					"Bacalhau com natas",
				},
				Dinner: []string{
					"Pizza",
					"Salada",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
