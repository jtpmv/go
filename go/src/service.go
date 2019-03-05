package main

import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	fmt.Println("Rendez-vous sur http://127.0.0.1:8000")

	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func HelloWorld(w  http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello World\n")
}
