package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	wc := map[string]int{}
	f := strings.Fields(s)
	for _, w := range f {
		if wc[w] == 0 {
			wc[w] = 1
		} else {
			wc[w]++
		}
	}

	return wc
}

func main() {
	WordCount("I am learning go")
}
