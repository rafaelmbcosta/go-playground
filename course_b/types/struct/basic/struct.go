package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// function with receiver (parameters first...)
// basically we can use method on product.
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	product1 := product{
		name:     "Pencil",
		price:    12.5,
		discount: 0.2,
	}

	product2 := product{"Notebook", 2040, 0.05}

	fmt.Println(product1, product1.priceWithDiscount())
	fmt.Println(product2, product2.priceWithDiscount())
}
