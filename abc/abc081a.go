package abc

import (
	"fmt"
)

// Abc081a is `ABC081A - Placing Marbles` answer
func Abc081a() {
	var s string
	fmt.Scanf("%s", &s)

	resultCount := 0
	if string(s[0]) == "1" {
		resultCount++
	}

	if string(s[1]) == "1" {
		resultCount++
	}

	if string(s[2]) == "1" {
		resultCount++
	}

	fmt.Println(resultCount)
}
