package main

import (
	"fmt"
)

func getType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return "other"
	}
}

func main() {
	fmt.Println(getType(123))
	fmt.Println(getType("hello"))
	fmt.Println(getType(true))
	fmt.Println(getType(make(chan interface{})))
	fmt.Println(getType(struct {
		nums []int
		name string
	}{}))
}
