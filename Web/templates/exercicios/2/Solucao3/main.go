package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hotel Abilio",
				Address: "42 Ferreira",
				City:    "Tomar",
				Zip:     "2240-424",
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
