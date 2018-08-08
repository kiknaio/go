package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	// Store to the memory address where int is stored
	var b *int = &a
	fmt.Println(b)

	// Differencing memory address
	fmt.Println(*b)

	// Change value of certain memory address
	*b = 42
	fmt.Println(a)
}
