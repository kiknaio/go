package main

import "fmt"

func main() {
	switch "George" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "George":
		fmt.Println("Whazaap George")
	default:
		fmt.Println("Who are you?")
	}
}
