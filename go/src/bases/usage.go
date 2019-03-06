package main

import(
	"fmt"
	"os"
	"log"
)

func Usage(){
	fmt.Println(os.Args[0], " <params>")
	os.Exit(1)
}

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	acpt := len(os.Args[1:])
	args := os.Args[1:]

	if acpt == 0 {
		Usage()
	}

	fmt.Println(args)
}
