package main

import "fmt"

func printApproved(approved ...string) {
	for i, guy := range approved {
		fmt.Printf("%d) %s\n", i+1, guy)
	}
}

func main() {
	app := []string{"Fabio", "Pedro", "Rui", "Miro"}
	printApproved(app...)
}
