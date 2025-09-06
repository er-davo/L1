package main

import "fmt"

func main() {
	a := 123
	b := 456

	fmt.Printf("a = %d, b = %d\n", a, b)

	a ^= b
	b ^= a
	a ^= b

	fmt.Printf("a = %d, b = %d\n", a, b)
}
