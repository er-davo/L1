package main

import "fmt"

func main() {
	var num, i, bit int64

	fmt.Printf("Num: ")
	fmt.Scan(&num)
	fmt.Printf("Bit: ")
	fmt.Scan(&i)
	fmt.Printf("Change to bit: ")
	fmt.Scan(&bit)

	a := int64(1) << i

	switch bit {
	case 1:
		fmt.Printf("Num: %b\nChanged: %b", num, num|a)
	case 0:
		fmt.Printf("Num: %b\nChanged: %b", num, num&^a)
	default:
		fmt.Println("Invalid bit value")
	}
}
