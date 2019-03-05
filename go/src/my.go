package	 main

import(
	"fmt"
	"log"
	"os"
	"bufio"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main(){

	f, err := os.Open("/tmp/fichier.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

