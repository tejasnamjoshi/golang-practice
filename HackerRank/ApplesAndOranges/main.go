package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var appleCount int32
	var orangeCount int32

	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			appleCount++
		}
	}

	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			orangeCount++
		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func main() {
	s := int32(2)
	t := int32(3)
	a := int32(1)
	b := int32(5)
	apples := []int32{2}
	oranges := []int32{-2}

	countApplesAndOranges(s, t, a, b, apples, oranges)
}
