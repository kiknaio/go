package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	secondSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, secondSlice...)

	fmt.Println(mySlice)
}
