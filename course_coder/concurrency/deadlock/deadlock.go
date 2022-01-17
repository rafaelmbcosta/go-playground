package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // blocking action
	time.Sleep(time.Second)
	fmt.Println("Only after read")
}

func main() {
	ch := make(chan int) // canal sem buffer

	go routine(ch)

	fmt.Println(<-ch) // blocking operation
	fmt.Println("Was read")

	fmt.Println(<-ch) // deadlock
	fmt.Println("end...")
}
