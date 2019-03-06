// jouer avec flag ( parser les arguments) 

package main

import(
	"fmt"
	"os"
)

func main(){

	cmd := os.Args[0]
	cptarg := len(os.Args[1:])

	fmt.Println(cmd)
	fmt.Println(cptarg)

}

