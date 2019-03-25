package main

import(
	"fmt"
	"log"
	"encoding/json"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

type Lien struct {
	titre	string
	url	string
}

type Page struct {
	Id	string `json:"id"`			// "-" si ne doit pas app
	Title	string `json:"title"`
}

func main(){
	b, err := json.Marshal(Page{"123","abcd"}); check(err)
	fmt.Println(string(b))
}


