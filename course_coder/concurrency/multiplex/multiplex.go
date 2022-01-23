package main

import (
	"fmt"
	"github/go-playground/html"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func join(data1, data2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(data1, c)
	go forward(data2, c)
	return c
}

// multiplexar seria misturar (mensagens) em um canal

func main() {
	c := join(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
