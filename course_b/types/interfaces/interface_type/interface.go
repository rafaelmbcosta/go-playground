package main

import "fmt"

type experiment struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	// virou budega, aceita qualquer coisa. (generic type)
	var anotherThing interface{}
	anotherThing = "Olá!"
	fmt.Println(anotherThing)
	anotherThing = false
	fmt.Println(anotherThing)
}
