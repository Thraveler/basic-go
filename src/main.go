package main

import "fmt"

func main() {
	// Constant declaration
	const piNumber float64 = 3.1416
	const name = "Daniel"

	fmt.Println(piNumber, name)

	// Variables
	age := 12
	var founds = 1500.34
	// var heigth float32

	fmt.Println("Age: ", age)
	fmt.Println("Founds: ", founds)
	// fmt.Println("Age: ", age)

	// Zero values
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}