package main

import(
	"fmt"
	"log"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

type Personne struct {
	firstname, lastname, email string
}

func main(){

	
}
