package main

import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	fmt.Println("Rendez-vous sur http://127.0.0.1:8888")

	http.HandleFunc("/jtpmv", index)
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func index(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "... Bienvenue sur JTPMV ... \n")
}
