package main

import (
	"log"
	"os"
	"text/template"
)

type region struct {
	Southern []hotel
	Central  []hotel
	Northern []hotel
}

type hotel struct {
	Name   string
	Adress string
	City   string
	Zip    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	Regions := region{
		Southern: []hotel{
			hotel{"Abilio", "Ferreira", "Tomar", "2240-424"},
			hotel{"Augusto", "Alenquer", "Tomar", "2310-324"},
		},
		Central: []hotel{
			hotel{"Hugo", "Ferreira", "Lisboa", "2240-424"},
			hotel{"Duarte", "Alenquer", "Lisboa", "2310-324"},
		},
		Northern: []hotel{
			hotel{"Andre", "Ferreira", "Porto", "2240-424"},
			hotel{"Fernando", "Alenquer", "Porto", "2310-324"},
		},
	}

	err := tpl.Execute(os.Stdout, Regions)
	if err != nil {
		log.Fatalln(err)
	}
}
