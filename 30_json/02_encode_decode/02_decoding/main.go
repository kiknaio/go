package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main()  {
	var p1 Person
	// Create a new reader and decode the JSON
	rdr := strings.NewReader(`{"First":"James","Last":"Bond","Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
