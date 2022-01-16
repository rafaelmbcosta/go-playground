package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getName() string {
	return p.name + " " + p.lastname
}

func (p *person) setFullName(fullname string) {
	parts := strings.Split(fullname, " ")
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := person{"Paulo", "Costa"}
	fmt.Println(p1, p1.getName())

	p1.setFullName("Raimundo Soldado")
	fmt.Println(p1, p1.getName())
}
