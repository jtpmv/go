package main

import (
	"fmt"
	"os"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	file, err := os.Open("file.go")
	check(err)

	data := make([]byte, 200)
	for count, err := file.Read(data) {

		check(err)
		fmt.Println(data[:count])
	}



	fmt.Println(data)
	fmt.Println("----")
	fmt.Print(count)

	/*
	for i := 0; i < count ; i++ {
		fmt.Print(data[i])
	}
	*/

	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// fmt.Println(data)
}
