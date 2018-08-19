package main

import "fmt"

func main() {
	students := []string{"George", "Luka", "Tata", "Gela-gocha"}
	// Delete Luka from students
	students = append(students[:1], students[2:]...)

	fmt.Println(students)
}
