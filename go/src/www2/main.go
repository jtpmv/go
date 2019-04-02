package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
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
		datas := strings.Split(rpath, "/")
		title := ""

		rule1, _ := regexp.MatchString("/favicon.ico", rpath)
		rule2, _ := regexp.MatchString("/toto", rpath)
		if !rule1 && !rule2 {

			fmt.Println(datas[2])
			title = datas[1]
		}

		data := Page{title, []string{"Go", "JavaScript", "Python"}}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe("localhost:8888", nil)
}
