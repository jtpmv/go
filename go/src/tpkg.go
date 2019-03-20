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

func main(){

	fmt.Println("Hello World !")
}

