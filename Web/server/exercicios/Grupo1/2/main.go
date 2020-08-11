package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Marcelo")
}

func a(w http.ResponseWriter, req *http.Request) {
	data := []string{"Abilio", "Marcelo", "Andre"}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", c)
	http.HandleFunc("/", a)
	http.ListenAndServe(":8080", nil)

}
