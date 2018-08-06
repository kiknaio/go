package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	// Print numbers in decimal, binary and hexidecimal
	for i:= 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t 0x%x \n", i, i, i)
	}
}
