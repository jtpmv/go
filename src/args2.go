package main

import (
	"fmt"
	"os"
)

func main(){
	a, sep := "",""
	for _, arg := range os.Args[1:] {
		a += sep + arg
		sep = " " 
	}
	fmt.Println(a)
}
