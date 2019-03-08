package main

import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	fmt.Println("Rendez-vous sur http://127.0.0.1:8888/jtpmv")

	http.HandleFunc("/jtpmv", srvWeb)
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func srvWeb(rep http.ResponseWriter, req *http.Request){
	io.WriteString(rep, "... Bienvenue sur JTPMV ... \n")
}
