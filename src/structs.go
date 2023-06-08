package src

import "fmt"

type car struct {
	make string
	year int
}

// Structs
func Structs() {
	myCar := car{make: "Tesla", year: 2023}

	fmt.Println(myCar)

	var mySecondCar car
	mySecondCar.make = "Suzuki"

	fmt.Println(mySecondCar)
}
