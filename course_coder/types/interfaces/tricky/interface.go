package main

import "fmt"

type sport interface {
	turnTurboOn()
}

type ferrari struct {
	model        string
	turboOn      bool
	currentSpeed int
}

func (f *ferrari) turnTurboOn() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnTurboOn()
	fmt.Println(car1)

	// var car2 sport = &ferrari{"F30", false, 0} ERROR

	var car2 sport = &ferrari{"F30", false, 0}
	car2.turnTurboOn()
	fmt.Println(car2)
}
