package main

import (
	"fmt"
	"time"
)

func twoThreeFour(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // linha tem que ser lido para o canal avançar

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int) // sem buffer
	go twoThreeFour(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo os 2 primeiro valores do canal
	fmt.Println(a, b)
	fmt.Println("B") // só chega nessa linha se os 2 dados do canal forem utilizados
	fmt.Println(<-c) // lendo o ultimo valor do canal

	// fmt.Println(<-c) // não tem mais ninguém, dá problema DEADLOCK

}
