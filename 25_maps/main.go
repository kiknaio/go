package main

import "fmt"

func main()  {
	// compose literal
	var students = map[string]int{}
	// create map with create
	var classes = make(map[string]string)

	// also compose literal
	courses := map[string]string{
		"CS101":	"Computer Science for beginners",
		"Math01":	"Mathematics for beginners",
	}

	students["George"] = 23
	students["Tata"] = 22

	classes["Math"] = "Mathematics"
	classes["CS"] = "Computer Science"

	fmt.Println(students)
	fmt.Println(classes)
	fmt.Println(courses)
}
