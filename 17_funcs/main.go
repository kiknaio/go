package main

import "fmt"

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

// Return multiple values
func multipleValues(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func main() {
	fullName := greet("George ", "Kiknadze")
	fmt.Println(fullName)

	fmt.Println(multipleValues("George ", "Kiknadze "))
}
