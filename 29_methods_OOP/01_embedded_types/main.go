package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

type MissionImpossibleTeam struct {
	Person
	Age int
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last: "Bond",
			Age: 20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last: "Moneypenny",
			Age: 18,
		},
		LicenseToKill: false,
	}

	// Override the types
	p3 := MissionImpossibleTeam{
		Person: Person{
			First: "Ethan",
			Last: "Hunt",
			Age: 40,
		},
		Age: 43,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)

	// Overriding example
	fmt.Println(p3.First, p3.Last, p3.Age)
}
