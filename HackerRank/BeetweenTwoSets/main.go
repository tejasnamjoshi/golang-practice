package main

import (
	"fmt"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	count := int32(0)
	l1 := len(a)
	l2 := len(b)

	for i := int32(1); i <= 100; i++ {
		factor := true
		for j := 0; j < l1; j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}

		if factor == false {
			continue
		}

		for k := 0; k < l2; k++ {
			if b[k]%i != 0 {
				factor = false
				break
			}
		}

		if factor == true {
			count++
		}
	}

	return count
}

func main() {
	a := []int32{1}
	b := []int32{100}

	count := getTotalX(a, b)

	fmt.Println(count)
}
