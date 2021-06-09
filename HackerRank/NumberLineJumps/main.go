package main

import (
	"fmt"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	k1 := x1
	k2 := x2

	if x1 > x2 && v1 > v2 || x2 > x1 && v2 > v1 {
		return "NO"
	}

	for i := 0; i < 20000; i++ {
		k1 += v1
		k2 += v2
		if k1 == k2 {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	x1 := int32(0)
	v1 := int32(2)
	x2 := int32(5)
	v2 := int32(3)

	result := kangaroo(x1, v1, x2, v2)
	fmt.Println(result)
}
