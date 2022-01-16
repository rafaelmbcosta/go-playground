package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalPrice() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	// fmt.Println("Total: ", total)
	return total
}

func main() {
	order1 := order{
		userID: 1,
		items: []item{
			{productID: 1, quantity: 2, price: 23.0},
			{productID: 2, quantity: 1, price: 27.0},
			{productID: 11, quantity: 4, price: 73.0},
		},
	}

	fmt.Println(order1, order1.totalPrice())
}
