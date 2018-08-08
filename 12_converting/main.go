package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := '🥕'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
	fmt.Println(string(129365))
}
