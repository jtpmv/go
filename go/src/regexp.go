package main

import(
	"fmt"
	"log"
	"os"
	"regexp"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func Usage(){
	fmt.Println("Usage:", os.Args[0], " <params>")
	os.Exit(1)
}

func main(){

	// RegExp
	re := regexp.MustCompile("t[a-zA-Z0-9]")

	if len(os.Args[1:]) == 0 { Usage() }

	for _,a := range os.Args[1:] {
		fmt.Println(re.MatchString(a))
	}
}

