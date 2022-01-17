package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags`
}

func main() {
	p1 := product{1, "notebook", 4500.23, []string{"What", "Ever"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json)) // slice of bytes

	var p2 product
	jsonString := `{"id":2,"name":"Pencil","pricing":23.4,"tags":["Oxente", "Rapaz"]}`
	json.Unmarshal([]byte(jsonString), &p2) // convert to byte and than sets p2
	fmt.Println(p2)
}
