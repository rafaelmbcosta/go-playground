package main

import (
	"fmt"
	"time"
)

func speak(person, text string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// Serial proccess
	// speak("Rafael", "Why so serious", 3)
	// speak("Mary", "Can only speak after you", 1)

	// go speak("John", "Hey....", 500)
	// go speak("Mary", "Wow !", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Finished !")

	go speak("John", "Understood!", 10)
	speak("Mary", "Congratulations", 5) // Main thread ends when Mary finishes.
}
