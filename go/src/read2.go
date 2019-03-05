package main

import(
	"fmt"
	"os"
	"io"
	"log"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

  file, err := os.Open("/tmp/fichier.txt")
  check(err)
  defer file.Close()

  buf := make([]byte, 1024)
  for {
	  n, err := file.Read(buf)
	  check(err)

	  if  err == io.EOF {
		  break
	  }

	  if n > 0 {
		  fmt.Print(string(buf[:n]))
	  }
  }
}

