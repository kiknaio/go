package main

// Pointer looks to memory address and changes the value of
// that memory address

// With * we are telling to modify the original value
// by getting the memory address
func zero(z *int)  {
	*z = 0
}

func main() {
	x := 5
	// & finds the memory address. Where it is stored
	zero(&x)
	println(x) // 0
}
