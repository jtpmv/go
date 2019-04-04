package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var cpt uint64
	var mutex = sync.Mutex{}

	for i := 0; i < 100; i++ {

		go func() {
			for {
				//time.Sleep(2 * time.Millisecond)

				mutex.Lock()

				cpt++
				mutex.Unlock()
			}
		}()

	}

	time.Sleep(time.Second)
	fmt.Println("Counter :", cpt)
}
