package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	s := []rune(str)

	// развернем всю строку а потом развернем слова
	// и тогда поменяется порядок слов но не буквы

	// развернем всю строку
	reverse(s, 0, len(s)-1)

	// развернем слова
	start := 0
	for end := 0; end < len(s); end++ {
		if s[end] == ' ' {
			reverse(s, start, end-1)
			start = end + 1
		}
	}

	// последнее слово
	reverse(s, start, len(s)-1)

	fmt.Println(string(s))
}
