package main
import(
	"fmt"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil { log.Fatal(e) }
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.Header.Get("User-Agent"))
		fmt.Fprintf(w, "Salut tout le monde" )
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Page de test " )
	})

	http.ListenAndServe("localhost:8888", nil)

}
