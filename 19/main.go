package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)

	s := []rune(str)

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	fmt.Println(string(s))
}
