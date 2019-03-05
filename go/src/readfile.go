package main

import (
	"fmt"
	"os"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


func main() {

	f, err := os.Open("file.txt")
	check(err)
	defer f.Close()

	// definir la taille du buffer
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		check(err)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}


