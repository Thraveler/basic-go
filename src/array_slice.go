package src

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func ArraySlice() {
	// Array -> inmutables
	array := [4]int{1, 2, 3}

	fmt.Println(array, cap(array), len(array))

	slices := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slices, cap(slices), len(slices))
	fmt.Println(slices[0])
	fmt.Println(slices[:3])
	fmt.Println(slices[3:])
	fmt.Println(slices[2:5])

	for i, _ := range slices {
		fmt.Println(i)
	}

}
