package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 10
	for counter < 20 {
		fmt.Println(counter)

		counter++
	}

	// for forever
	for {
		fmt.Println("Until the end")
	}
}
