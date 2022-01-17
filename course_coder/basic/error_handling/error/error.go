package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	fmt.Println(f, err)
}
