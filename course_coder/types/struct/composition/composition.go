package main

import "fmt"

type car struct {
	name  string
	speed int
}

type ferrari struct {
	car     // anonymous fields
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.speed = 40
	f.turboOn = true
	fmt.Printf("Ferrari %s have its turbo on? %v", f.name, f.turboOn)
}
