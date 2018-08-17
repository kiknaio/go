package main

import "fmt"

func main() {
	mySlice := []string{"ab", "cd", "ef", "gh", "ij"}

	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	// Get the character code from ASCII and UTF8
	fmt.Println("myString"[1])
}
