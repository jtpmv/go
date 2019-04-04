package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
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

		rpath := r.URL.Path
		rule1, _ := regexp.MatchString("/favicon.ico", rpath)
		rule2, _ := regexp.MatchString("/toto", rpath)
		title := ""

		if !rule1 && !rule2 {
			queries := r.URL.Query()
			title = queries.Get("nom")
			age := queries.Get("age")

			fmt.Println(queries.Get("age"))
		}

		data := Page{title, []string{"Go", "JavaScript", "Python"}}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe("localhost:8888", nil)
}
