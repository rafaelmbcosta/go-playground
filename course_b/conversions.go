package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// int to float
	fmt.Println(x / float64(y))
	fmt.Println(x + float64(y))

	// int to string
	fmt.Println("Test " + strconv.Itoa(123))
	fmt.Println("Test", 123)

	// string p/ int

	num, _ := strconv.Atoi("123")
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))

	// string p/ bool

	bool, _ := strconv.ParseBool("false")
	fmt.Println(bool)
	fmt.Println(reflect.TypeOf(bool))

	// string p/ float

	float, _ := strconv.ParseFloat("49.09", 64)
	fmt.Println(float)
	fmt.Println(reflect.TypeOf(float))

}
