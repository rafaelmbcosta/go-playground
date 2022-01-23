package main

import (
	"fmt"
	"github/go-playground/html"
	"time"
)

func fastest(url1, url2, url3 string) string {
	c1 := html.Titulo(url1) // this function returns one channel
	c2 := html.Titulo(url2) // this function returns one channel
	c3 := html.Titulo(url3) // this function returns one channel

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(400 * time.Millisecond):
		return "All is lost"
		// default:
		// 	return "No response yet !"
	}

}

func main() {
	winner := fastest(
		"https://www.youtube.com.br",
		"https://www.google.com",
		"https://www.amazon.com.br",
	)
	fmt.Println(winner) // with 400 ms, google wins
}
