package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("Executed !") // not executed
	c <- 5
	c <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch) // 1
	// It only blocks when buffer is full
	fmt.Println(<-ch) // 1

}
