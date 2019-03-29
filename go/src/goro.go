package main

import(
	"fmt"
	"time"
)

func roll(name string, loops int,sec int){
	for i:=1; i<=loops; i++ {
		fmt.Printf("%s -----> %d \n", name, i)
		time.Sleep(time.Second * time.Duration(sec))
	}
}

func main(){
	//fmt.Println("Debut")
	//time.Sleep(3*time.Second)
	//fmt.Println("Fin")
	go roll("Un",4,1)
	go roll("Deux",4,1)
	time.Sleep(time.Second * 6)
	fmt.Println("Fin")


}

