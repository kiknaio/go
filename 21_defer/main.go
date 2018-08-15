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

func main()  {
	defer world()
	test()
	hello()
}
