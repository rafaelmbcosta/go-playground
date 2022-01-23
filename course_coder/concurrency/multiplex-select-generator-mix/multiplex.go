package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}() // invoca a funcão automaticamente

	return c
}

func join(data1, data2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-data1:
				c <- s
			case s := <-data2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := join(speak("Zanata"), speak("Ricardo"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
