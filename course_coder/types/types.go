package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Integers
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is", reflect.TypeOf(32000))
	fmt.Println("Integer literal negative is", reflect.TypeOf(-32000))

	// uint8 uint16 uint32 uint64

	// byte = uint8
	var b byte = 255
	fmt.Println("O byte é: ", reflect.TypeOf(b))

	// max Int
	// maxInt := math.MaxInt64
	// fmt.Println("Max Int = ", maxInt)

	// Rune table Unicode
	var rn rune = 'a'
	fmt.Println("O rune é: ", reflect.TypeOf(rn))
	fmt.Println(rn)

	// floats
	var x float32 = 49.99
	y := 49.99
	fmt.Println("Tipo x é ", reflect.TypeOf(x))
	fmt.Println("Valor do Tipo Literal é ", reflect.TypeOf(49.99))
	fmt.Println(reflect.TypeOf(y))

	// boolean
	bo := false
	fmt.Println("Tipo de bo é", reflect.TypeOf(bo))

	// strings
	s1 := "Olá me llamo Rafael"
	fmt.Println(s1 + "!")
	fmt.Println(len(s1))

	// char é um uint32
	// não existe o tipo char.
	char := 'a'
	fmt.Println(reflect.TypeOf(char))

}
