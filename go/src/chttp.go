package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	rep, err := http.Get("http://example.com")
	check(err)
	//fmt.Println(rep)

	body, err := ioutil.ReadAll(rep.Body)
	defer rep.Body.Close()
	fmt.Println(string(body))

}
