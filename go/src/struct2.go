package main

import("fmt")

type Personne interface {
	getName() string
}

type Employe struct {
	id	int
	firstname, lastname, email string
}

// En Go, il est possible de déclarer 
// des structures et des fonctions membres de ces structures
func (e Employe) getName() string {
	return e.firstname
}

func foo(p Personne) {
	fmt.Println(p.getName)
}



func main(){

	ee := Employe{1, "Johnny", "Durand", "johnny.durand@gmail.com" }
	foo(ee)
}
