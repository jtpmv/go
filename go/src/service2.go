package main

import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	fmt.Println("Rendez-vous sur http://127.0.0.1:8800/toto")

	http.HandleFunc("/toto", index)
	http.ListenAndServe("127.0.0.1:8800", nil)
}

func index(w  http.ResponseWriter, req *http.Request){
	io.WriteString(w, "... A venir ... \n")
}
