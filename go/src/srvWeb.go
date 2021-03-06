package main

import(
	"fmt"
	"log"
	"io"
	"net/http"
)

// struct page 

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func requete(uri, content string){

	http.HandleFunc(uri, func(response http.ResponseWriter, request *http.Request){
		io.WriteString(response, content)
		log.Println(uri, "---",  response)
	})
}

func main(){

	fmt.Println("Rendez-vous sur http://127.0.0.1:8888")

	requete("/", "... SRV GO - Bienvenue chez moi ;-) ...\n")
	requete("/jtpmv", "... Bienvenue sur JTPMV ...\n")
	requete("/titi", titi)

	http.ListenAndServe("127.0.0.1:8888", nil)
}





var titi string = `
<html>
<head>
  <style>
  body { background-color: black; color: white; } 
  </style>
</head>
<body>
	<h2>Titre</h2>
	<h3>Sous-titre</h3>
	Test d'interpretation ... <br>
	<b>Autres ...</b>

<body>
</html>`


