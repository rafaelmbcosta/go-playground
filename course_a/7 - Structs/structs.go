package main

import "fmt"

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number int8
}

// animal := { oi name }
func main() {
	var u user
	u.name = "Rafael"
	u.age = 26
	// fmt.Println(u)
	fmt.Println(u.name)

	address1 := address{"Baker Street", 122}

	u2 := user{"Paulo", 20, address1}
	fmt.Println(u2)

	u3 := user{name: "Paulo"}
	fmt.Println(u3)

}
