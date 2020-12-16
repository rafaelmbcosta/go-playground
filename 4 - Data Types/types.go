package main

import "fmt"

func main() {
	// int8 (byte), int16, int32 (rune), int64, int, uint
	// var number int64 = 1000
	var number int = 1000000000
	negativeNumber := -12
	fmt.Println(number)
	fmt.Println(negativeNumber)

	// var number2 uint32 = -11 Error
	var number2 uint32 = 11
	fmt.Println(number2)

	// float32, float64
	var float1 float32 = 12.2
	fmt.Println(float1)

	floatAuto := 15.5
	fmt.Println(floatAuto)
}
