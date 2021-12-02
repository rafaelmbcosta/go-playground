package main

import "fmt"

func main() {
	x := 3.141517
	fmt.Println("Valor de x", x)
	// fmt.Println("Valor de x " + x)
	xs := fmt.Sprint(x)
	fmt.Println("(sprint) Valor de x " + xs)
	fmt.Printf("(interpolation) Valor de x é %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "ops"

	fmt.Printf("\n%d %f %t %s,", a, b, c, d)
	fmt.Printf("\n%d %.1f %t %s \n", a, b, c, d)
	fmt.Println(a, " ", b, " ", c)
}
