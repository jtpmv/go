package main

import "fmt"

func fp(a *[5]int) { fmt.Println(a) }

func main(){

	for i := 0; i < 5; i++ {
		fp(&[5]int{i, i*i, i*i*i})
	}


}
