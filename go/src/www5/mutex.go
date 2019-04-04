package main

// synchronisation des channels : les mutex
import (
	"fmt"
	"time"
)

func main() {

	//timer := time.NewTimer(4 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick à :", t)
		}
	}()

	time.Sleep(6 * time.Second)
	ticker.Stop()
	time.Sleep(2 * time.Second)

	fmt.Println("Fin de programme")

}
