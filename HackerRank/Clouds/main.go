package main

import (
	"fmt"
)

func jumpingOnClouds(c []int32) int32 {
	var n int32 = 0
	l := len(c)

	for i := 0; i < l-1; {
		if i+2 == l || c[i+2] == 1 {
			i++
			n++
		} else {
			i += 2
			n++
		}
	}

	return n
}

func main() {
	c := []int32{0, 0, 0, 0, 1, 0}
	result := jumpingOnClouds(c)
	fmt.Println(result)
}
