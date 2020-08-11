package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Marcelo")
}

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Website")
}
func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", c)
	http.HandleFunc("/", a)
	http.ListenAndServe(":8080", nil)
}
