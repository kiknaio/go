package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 90}, "Computer Science"}

	fmt.Printf("%s is %d years old and he knows %s", mark.name, mark.age, mark.speciality)
}
