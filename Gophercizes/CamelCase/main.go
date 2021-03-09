package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
		// if ch < 'Z' && ch > 'A' {
		// 	answer++
		// }
		// if ch < 91 && ch > 63 {
		// 	answer++
		// }
	}

	fmt.Println(answer)

}
