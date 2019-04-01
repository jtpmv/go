package main

import(
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

type Page struct {
	Title string
	Menu []string
}

func main(){

	tmpl, err := template.ParseGlob("./templates/*.html"); check(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		data := Page{"Langages", []string{"Go","JavaScript","Python"}}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	//fileserver := http.FileServer(http.Dir("./assets"))
	//http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
