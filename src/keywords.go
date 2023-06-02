package main

import "fmt"

func main() {
	defer fmt.Println("This will be executed at the end!")

	counter := 0
	for {
		counter++

		if counter%2 == 0 {
			continue
		}

		if counter >= 10 {
			break
		}

		fmt.Println(counter)
	}
}
