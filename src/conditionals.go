package src

import (
	"fmt"
	"log"
	"strconv"
)

func Conditionals() {
	for i := 1; i <= 4; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Is even")
		} else {
			fmt.Println(i, " Is odd")
		}
	}

	result, err := strconv.Atoi("23")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Result: ", result)
	}

}
