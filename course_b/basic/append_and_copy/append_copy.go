package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// Não permite por não ser slice.
	// slice1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)

	slice2 := make([]int, 3)
	copy(slice2, slice1)

	fmt.Println(array1, slice1, slice2)
}
