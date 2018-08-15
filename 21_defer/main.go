package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world()  {
	fmt.Println("world")
}

func test() {
	fmt.Println("Test")
}

// Defer will run after everything is done inside of the main
func main()  {
	defer world()
	test()
	hello()
}
