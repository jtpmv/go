package main

import (
	"fmt"
	"log"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func roll(name string, loops int, sec int, c *chan bool) {
	for i := 0; i < loops; i++ {
		fmt.Printf("%s ---> %d \n", name, i)
		time.Sleep(time.Duration(sec) * time.Second)
	}
	*c <- true
}

func main() {
	interrupt := make(chan bool)
	go roll("R", 4, 1, &interrupt)
	go roll("RR2", 5, 2, &interrupt)
	go roll("RRR3", 5, 1, &interrupt)

	for i := 1; i <= 3; i++ {
		<-interrupt
	}

	fmt.Println("Fin de programme")
}
