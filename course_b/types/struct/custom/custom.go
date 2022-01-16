package main

import "fmt"

type note float64

func (n note) between(first, final float64) bool {
	return float64(n) >= first && float64(n) <= final
}

func noteForConcept(n note) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else {
		return "B"
	}
}

func main() {
	fmt.Println(noteForConcept(6.5))
}
