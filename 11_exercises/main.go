package main

import "fmt"

func main() {
	// Ask names
	var name string
	fmt.Print("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Hello " + name)

	//	Reminders
	var smallNumber int
	var bigNumber int

	fmt.Print("Please provide the small number ")
	fmt.Scan(&smallNumber)
	fmt.Print("Please provide the big number ")
	fmt.Scan(&bigNumber)

	answer := bigNumber % smallNumber
	fmt.Printf("%d divided by %d and the reminder is %d\n", bigNumber, smallNumber, answer)

	// Slice of bytes
	fmt.Println([]byte("George"))
}
