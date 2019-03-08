package main

//func EscapeString(s string) string
//func UnescapeString(s string) string
// ./clt | sed -r 's/<\/?[a-zA-Z]+[0-9]*>//g'


import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/denisbrodbeck/striphtmltags"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	page, err := http.Get("http://127.0.0.1:8888/titi")
	check(err)

	body, err := ioutil.ReadAll(page.Body)
	defer page.Body.Close()
	fmt.Println(string(body))
	fmt.Println("-----")

	fmt.Println(striphtmltags.StripTags(string(body)))
}
