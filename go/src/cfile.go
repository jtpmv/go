package main

import(
	"fmt"
	"os"
	"io/ioutil"
)


func check(e error){
	if e != nil {
		fmt.Println(e)
	}
}

func main(){

	fmt.Println("Creation du fichier /tmp/data1.txt --- 1ere façon: ioutil")

	s1 := []byte("Hello World! \n Create file ;-)\n")
	f4 := ioutil.WriteFile("/tmp/data1.txt", s1, 0644)
	check(f4)


	fmt.Println("Lecture du fichier /tmp/data1.txt ")
	f5, err := os.Open("/tmp/data1.txt")
	check(err)

	buf2 := make([]byte, 100)
	for {
		cpt, err := f5.Read(buf2)
		check(err)
		if cpt > 0 {
			fmt.Print(string(buf2[:cpt]))
		}
	}
}
