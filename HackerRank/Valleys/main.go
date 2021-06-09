package main

import (
	"fmt"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	var v int32 = 0
	pos := 0
	s := true

	for i := int32(0); i < steps; i++ {
		if string(path[i]) == "U" {
			pos++
		} else {
			pos--
		}

		if s && pos == -1 {
			v++
			s = false
		}

		s = pos == 0
	}

	return v
}

func main() {
	var steps int32 = 8
	path := "UDDDUDUU"
	result := countingValleys(steps, path)

	fmt.Println(result)
}
