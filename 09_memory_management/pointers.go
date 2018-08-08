package main

// Pointer looks to memory address and changes the value of
// that memory address
func zero(z *int)  {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	println(x) // 0
}
