package	main

import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	page1 := func(w http.ResponseWriter, r  *http.Request){
			io.WriteString(w, "Hello . 1 !\n")
	}

	page2 := func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello .. 2 !!\n")
	}

	page3 := func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello ... 3 !!!\n")
	}

	http.HandleFunc("/", page1)
	http.HandleFunc("/toto", page2)
	http.HandleFunc("/titi", page3)

	//log.Fatal(http.ListenAndServe("localhost:8080", nil))
	fmt.Println("http://localhost:8080")
	fmt.Println("http://localhost:8080/toto")
	fmt.Println("http://localhost:8080/titi")
	http.ListenAndServe("localhost:8080", nil)

}


