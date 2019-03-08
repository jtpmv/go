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

	rep, err := http.Get("http://127.0.0.1:8888/titi")
	check(err)
	//fmt.Println(rep)

	body, err := ioutil.ReadAll(rep.Body)
	defer rep.Body.Close()
	fmt.Println(string(body))

}
