package main

import "fmt"

func main() {
	// Initialize empty array
	var array [5]int

	// Fixed size array literal
	secondArray := [5]int{10, 20, 30, 40, 50}

	// Not-fixed size array
	thirdArray := [...]int{10, 20, 30}

	// Declare an integer array of five elements.
	// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	exampleArray := [5]int{1: 10, 2: 20}
	exampleArray[3] = 15

	// Check copy of array
	copyArray := exampleArray
	copyArray[1] = 15

	fmt.Println(exampleArray)
	fmt.Println(copyArray)

	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	pointerArray := [5]*int{0: new(int), 1: new(int)}

	// Assign values to index 0 and 1.
	*pointerArray[0] = 10
	*pointerArray[1] = 20


	fmt.Println(array, secondArray, thirdArray, exampleArray)
	fmt.Println(pointerArray)
}
