package main 

import(
	"fmt"
	"log"
	"net/http"
)

func ExampleFileServer() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}

func main(){
	fmt.Println("http://localhost:8080  => /usr/share/doc")
	ExampleFileServer()

}
