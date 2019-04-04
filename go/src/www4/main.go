package main

import (
	"fmt"
	"log"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Codeur struct {
	Nom string
	Age uint8
	Skills []string `json:"Langages"`
	// OR
	//Skills interface{} `json:"Langages"`  // plus utilisable en tant que array or slice (facon string)
}

func main() {

	//alex := Codeur{"Alex",26, []string{"Go", "Javascript", "Python"}}
	//js, _ := json.Marshal(alex)

	//fmt.Println(string(js))
	//fmt.Println(alex.Skills)

	co := &Codeur{}
	steph := `{"Nom":"Alex","Age":26,"Langages":["Go", "Javascript", "Python"]}`
	json.Unmarshal([]byte(steph), co)

	fmt.Println(co)
	//fmt.Println(co.Skills[2])
	fmt.Println(co.Skills)

}
