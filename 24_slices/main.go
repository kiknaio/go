package main

import "fmt"

func main()  {
	// Create a new slice of integers with the capacity
	// of 5 and use 0 (underlying array has length of 5)
	mySlice := make([]int, 0, 5)

	fmt.Println("-----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	// Get capacity
	fmt.Println(cap(mySlice))
	fmt.Println("-----------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
