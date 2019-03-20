package main

//func EscapeString(s string) string
//func UnescapeString(s string) string
// ./clt | sed -r 's/<\/?[a-zA-Z]+[0-9]*>//g'


import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"regexp"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	//uri := "http://127.0.0.1:8888/titi"
	//uri := "https://www.zdnet.fr/actualites/5g-les-premieres-frequences-s-arrachent-a-188-millions-d-euros-en-autriche-39881773.htm"
	uri := "https://linuxfr.org/users/nono/journaux/go-et-rust-lequel-est-le-remplacant-du-c"

	page, err := http.Get(uri)
	check(err)

	//re := regexp.MustCompile(`<-?!?/?[a-zA-Z]+[-a-zA-Z0-9\=\"\'\/\?\_\-\#\.\ \t\&\;]*>`)
	re := regexp.MustCompile(`<-?!?/?[-a-zA-Z]+.*>`)

	body, err := ioutil.ReadAll(page.Body)
	defer page.Body.Close()


	fmt.Println(re.ReplaceAllString(string(body), ""))

	//fmt.Println("-----")
	//fmt.Println((string(body)))
}
