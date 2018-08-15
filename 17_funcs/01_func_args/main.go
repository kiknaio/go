package main

import "fmt"

func average(sf ...float64) float64  {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	n := average(43, 56, 21, 63, 63, 33, 12, 89, 42)
	fmt.Println(n)

	// Only one way to define functions inside the function
	sayHello := func() {
		fmt.Println("Hello")
	}

	sayHello()
}