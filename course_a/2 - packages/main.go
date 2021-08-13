// vai saber se o arquivo é executavel
package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Importing main file")
	auxiliar.Writing()
	error := checkmail.ValidateFormat("rafael")
	fmt.Println(error)
}
