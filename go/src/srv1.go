package main

import(
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Personne struct {
	Firstname, Lastname, Email string
}

type Page struct {
	Titre string
	Content string
}

func main(){

	tmpl, _ := template.ParseGlob("./www/*.html")   // {{ define "PageName" }} ... {{ end }} 
	data := Page{"Ma page", "<h1>Contenu de ma page</h1>" }

	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request){
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/jtpmv" , func(w http.ResponseWriter, r *http.Request){
		data1 := Page{"Ma page Index2", "Contenu de data pour ma page 2" }
		tmpl.ExecuteTemplate(w, "index2", data1)
	})

	http.HandleFunc("/jtpmvrc" , func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Bienvenue chez jtpmvrc !")
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe("localhost:8888", nil)
}
