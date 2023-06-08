package src

import "fmt"

func Maps() {
	var persons = make(map[string]int)

	persons["Daniel"] = 28
	persons["Juan"] = 35

	fmt.Println(persons)

	for key, value := range persons {
		fmt.Println(key, value)
	}

	key, okkk := persons["Josedeodo"]

	fmt.Println(key, okkk)

	planets := map[string]string{
		"earth": "Planeta Tierra",
		"mart":  "Planeta Rojo",
	}

	for _, planet := range planets {
		fmt.Println(planet)
	}
}
