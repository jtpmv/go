package main

import (
	"fmt"
	"os"
)

func main(){
	a, sep := "",""
	for i := 1 ; i < len(os.Args); i++ {
		a += sep + os.Args[i]
		sep = " " 
	}
	fmt.Println(a)
}
