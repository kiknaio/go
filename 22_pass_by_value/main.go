package main

import "fmt"

func changeMe(z *int) {
	var a uint

	a = -2

	fmt.Println(a)

	fmt.Println(z) // 0xc420016090
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z) // 0xc420016090
	fmt.Println(*z) // 24
}

func main() {
	age := 44
	fmt.Println(&age) // 0xc420016090

	changeMe(&age)

	fmt.Println(&age) // 0xc420016090
	fmt.Println(age) // 24
}
