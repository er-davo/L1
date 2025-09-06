package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, v := range s {
		if !set[v] {
			set[v] = true
		}
	}

	for k := range set {
		fmt.Printf("%s ", k)
	}
}
