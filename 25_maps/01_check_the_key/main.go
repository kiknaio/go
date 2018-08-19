package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno",
	}

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
}
