package main

import "fmt"

func main() {

	option := 1

	switch option {
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("Default")
	}

	switch {
	case option == 1:
		fmt.Println("1")
	default:
		fmt.Println("Default")
	}
}
