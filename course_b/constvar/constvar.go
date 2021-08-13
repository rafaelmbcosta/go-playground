package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo inferido pelo compilador

	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferencia é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	e, f, g := 5, 6, 7

	fmt.Println(a, b, c, d, e, f, g)
}
