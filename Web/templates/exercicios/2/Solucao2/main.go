package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hotel Abilio",
			Address: "42 Ferreira",
			City:    "Tomar",
			Zip:     "2240-424",
			Region:  "southern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
