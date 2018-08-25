package main

import "fmt"

type person struct {
	firstname string
	lastname string
	age int
}

func main()  {
	p1 := person{firstname: "George", lastname: "Kiknadze", age: 23}
	p2 := person{firstname: "Luka", lastname: "Kiknadze", age: 19}

	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
