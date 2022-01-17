package main

import "fmt"

func main() {
	ch := make(chan int, 1) // tipo do canal e o buffer
	ch <- 1                 // passando para o canal o valor 1 (escrita)
	<-ch                    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
