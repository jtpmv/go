package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func usage(){
	fmt.Println("Usage: ", os.Args[0], "[url]")
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	//if os.Args[1] { usage() }


	re := regexp.MustCompile("http.*")

	for _, a := range os.Args[1:] {

		if re.MatchString(a) {
			fmt.Println(a)
		}
	}
}

