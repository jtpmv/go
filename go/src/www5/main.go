package main

import "log"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	db, err := bolt.Open("dbtest.db", 0775, nil)
	defer db.Close()

	check(err)

}
