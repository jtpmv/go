package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Client struct {
	Nom   string
	Id uint `json:"Numéro_Client"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rpath := r.URL.Path
		rule1, _ := regexp.MatchString("/favicon.ico", rpath)

		if !rule1 {
			queries := r.URL.Query()
			client := Client{}
			client.Nom = queries.Get("nom")

			id, _ := strconv.Atoi(queries.Get("id"))
			client.Id = uint(id)

			fmt.Println(client)

			//js, _ := json.Marshal(client)
			//fmt.Fprintf(w, string(js)  )
			json.NewEncoder(w).Encode(client)
		}
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe("localhost:8888", nil)
}
