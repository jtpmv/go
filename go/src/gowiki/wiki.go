package main

import(
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// methode de sauvegarde de page 
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// methode de chargement de page 
func loadPage(title string) *Page {
	filename :=  title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}


