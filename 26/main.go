package main

import (
	"fmt"
	"strings"
)

func isUniceBytes(s []byte) bool {
	// быстрее чем через мапу, но для ограниченного набора символов (all ASCII characters)
	var seen [256]bool

	for _, v := range s {
		if seen[v] {
			return false
		}
		seen[v] = true
	}

	return true
}

func isUniceRunes(s []rune) bool {
	// чуть больше памяти
	seen := make(map[rune]bool)

	for _, v := range s {
		if seen[v] {
			return false
		} else {
			seen[v] = true
		}
	}

	return true
}

func main() {
	var s string

	fmt.Scan(&s)

	s = strings.ToLower(s)

	fmt.Printf("Is unice bytes: %v\nIs unice runes: %v\n",
		isUniceBytes([]byte(s)), isUniceRunes([]rune(s)))
}
