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
		n, _ := f.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		} else {
			break
		}
	}
}


