package main

import(
	"fmt"
	"log"
	"os"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func data(filename string) string {
	f, _ := os.Open(filename)
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil { os.Exit(0) }

		if n > 0 {
			fmt.Print(string(buf[:n]))
		}

	}
}

func main(){
	data("file.txt")
}
