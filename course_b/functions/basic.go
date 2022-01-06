package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s - %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s \n", p1, p2)
}

func f5() (string, string) {
	return "retorno 1", "retorno 2"
}

func main() {
	f1()
	f2("oi", "curioso")
	fmt.Println(f3())
	fmt.Println(f4("oi", "curioso"))
	r1, r2 := f5()
	fmt.Println(r1, r2, "Retornos de F5")
}
