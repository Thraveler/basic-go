package main

import "fmt"

// Simple function
func helloWorld () {
	fmt.Println("Hello World!")
}

// Function with parameters
func greetingTo(name string, lastname string) {
	fmt.Printf("Hello %s %s\n", name, lastname)
}

// Using a short way to declarate parameters and returning a value
func addition(x, y int) int {
	return x + y
}

// Returning multiple values
func multiplyByTwo(x int) (int, int) {
	return x, x * 2
}

func main() {
	helloWorld()
	greetingTo("Jane", "Doe")

	fmt.Println(addition(5, 4))

	originalValue, result := multiplyByTwo(2)
	fmt.Printf("Original value %d, result multiply by 2: %d", originalValue, result)

	// Taking only one return value
	_, multiplyByTwo := multiplyByTwo(10)
	fmt.Println(multiplyByTwo)
}