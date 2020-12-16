package main

import "fmt"

type person struct {
	name string
	age  int8
}

// just type the other struct name
type student struct {
	person
	university string
}

func main() {
	p1 := person{name: "Zora"}
	p2 := student{p1, "Berilo"}
	fmt.Println(p2.name)
}
