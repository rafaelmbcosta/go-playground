package main

import "fmt"

// copy value
func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * get the value
	*n++
}

func main() {
	x := 1
	fmt.Println(x)
	inc1(x)
	fmt.Println(x)
	inc2(&x)
	fmt.Println(x)
}
