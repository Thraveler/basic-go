package main

import "fmt"

func main () {
	hello := "Hello"
	world := "World"
	name := "Peter"
	age := 32

	// Print text and variables adding a line break at the end
	fmt.Println("Some text", hello, world)

	// Allows you to format the text as you want knowing data type
	fmt.Printf("My name is %s and my age is %d.\n", name, age)

	// Allows you to format the text as you want without knowing data type
	fmt.Printf("My name is %v and my age is %v.\n", name, age)

	// Create a new string with format
	text := fmt.Sprintf("My name is %v and my age is %v.\n", name, age)
	fmt.Println(text)

	fmt.Printf("Name type is: %T \nAge type is %T.\n", name, age)

}