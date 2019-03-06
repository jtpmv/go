package main

import(
	"fmt"
	"os"
	"log"
)

func Usage(){
	fmt.Println("Usage: ", os.Args[0], " <params>")
	os.Exit(1)
}

func Help(){
	fmt.Println("No help !")
}

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	acpt := len(os.Args[1:])
	if acpt == 0 {
		Usage()
	}

	args := os.Args[1:] 
	for _, a := range args {
		if a == "h" { Help() }
	}
}
