package main

type sportive interface {
	turboOn()
}

type deluxe interface {
	shine()
}

type sportiveDeluxe interface {
	sportive
	deluxe
}
