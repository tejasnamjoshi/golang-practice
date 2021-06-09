package main

import (
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32 = 0

	for i := 0; i < len(ar); i++ {
		if ar[i] == -1 {
			continue
		}
		for j := i + 1; j < len(ar); j++ {
			if ar[i] == ar[j] {
				pairs += 1
				ar[j] = -1
				break
			}
		}
	}

	return pairs
}

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	var n int32 = 9
	result := sockMerchant(n, ar)

	fmt.Println(result)
}
