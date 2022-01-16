package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

// implicit declaration
func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - $ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	person1 := person{"Rafael", "Costa"}
	fmt.Println(person1.toString())
	print(person1)

	thing := product{"Tapioca", 23.4}
	fmt.Println(thing.toString())
	print(thing)
}
