package main

import "fmt"

func f(msg string) {
	fmt.Println(msg)
}

func main() {
	go f("go routine")
	f("function")
}
