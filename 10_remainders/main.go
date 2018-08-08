package main

import "fmt"

func main() {
	x := 13 % 3

	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Event")
	}

	for i := 0; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Event")
		}
	}
}
