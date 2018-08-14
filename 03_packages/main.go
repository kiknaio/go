package main

import (
	"zero_to_hero/03_packages/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.Reverse("Kiknadze"))
	fmt.Println(utils.Reverse(utils.Name))
}