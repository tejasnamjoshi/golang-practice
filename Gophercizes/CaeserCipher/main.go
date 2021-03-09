package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	// var k int
	fmt.Scanf("%s\n", &input)
	// fmt.Scanf("%d\n", &k)
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ret := ""
	for _, ch := range input {
		if strings.IndexRune(alphabetLower, ch) >= 0 {
			ret += string(rotate(ch, 2, []rune(alphabetLower)))
		} else if strings.IndexRune(alphabetUpper, ch) >= 0 {
			ret += string(rotate(ch, 2, []rune(alphabetUpper)))
		} else {
			ret += string(ch)
		}
	}
	fmt.Println(string(ret))
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}

	idx = (idx + delta) % len(key)

	return key[idx]

	// idx := -1
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }

	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx >= len(key) {
	// 		idx = 0
	// 	}
	// }

	// return key[idx]
}
