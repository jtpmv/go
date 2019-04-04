package main

import (
	"fmt"
	"time"
)

//func roll(name string, loops int, sec int, c *chan bool) {
//	for i := 0; i < loops; i++ {
//		fmt.Printf("%s ---> %d \n", name, i)
//		time.Sleep(time.Duration(sec) * time.Second)
//	}
//	*c <- true
//}

func main() {
	c := make(chan bool)

	go func() {
		time.Sleep(4 * time.Second) // traitement durée 4 sec.
		c <- true
	}()

	select {
	case b1 := <-c:
		fmt.Println(b1)
	case <-time.After(2 * time.Second):
		fmt.Println("Ne répond pas")
	}
	fmt.Println("Fin de programme")
}
