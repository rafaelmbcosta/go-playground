package main

import "fmt"

func main() {
	funcsByLetter := map[string]map[string]float64{
		"G": {
			"Gabriel": 14554.93,
			"Gustavo": 555.5,
		},
		"F": {
			"Fabio":  12000.9,
			"Fulano": 3333.0,
		},
	}

	fmt.Println(funcsByLetter)
	fmt.Println(funcsByLetter["G"]["Gabriel"])
	delete(funcsByLetter, "F")
	fmt.Println(funcsByLetter)
}
