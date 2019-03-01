package main

import(
	"fmt"
	"io/ioutil"
)


func check(e error){
	if e != nil {
		fmt.Println(e)
	}
}


func main(){

	s1 := []byte("Hello World! \n Create file ;-)\n")
	f := ioutil.WriteFile("/tmp/data1.txt", s1, 0644)
	check(f)


}
