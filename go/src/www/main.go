package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Page struct {
	Title   string
	Content []string
}

func main() {

	tmpl, err := template.ParseGlob("./templates/*.html")
	check(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := Page{"Langages", []string{"Go", "JavaScript", "Python"}}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe("localhost:8888", nil)
}
