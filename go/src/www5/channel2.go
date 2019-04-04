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
	msg1 := make(chan string)
	msg2 := make(chan string)

	go func() {

		msg1 <- "Un ..."
		time.Sleep(2 * time.Second)
		msg2 <- "Deux ..."

	}()

	for i := 1; i <= 2; i++ {
		select {
		case m1 := <-msg1:
			fmt.Println(m1)
		case m2 := <-msg2:
			fmt.Println(m2)
		}
	}
}
