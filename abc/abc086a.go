package abc

import "fmt"

// Abc086a is `ABC086A - Product` answer
func Abc086a() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	multiplication := a * b
	if multiplication%2 == 0 {
		fmt.Println("Even")
		return
	}

	fmt.Println("Odd")
}
