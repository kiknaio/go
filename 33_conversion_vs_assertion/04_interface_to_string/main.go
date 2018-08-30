package main

import "fmt"

func main() {
	var name interface{} = "Tbilisi"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n %s\n", str, name)
	} else {
		fmt.Println("Value is not a string")
	}
}
