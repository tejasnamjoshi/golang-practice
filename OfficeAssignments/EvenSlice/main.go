package main

import "fmt"

func main() {
	var is []int
	var i int
	fmt.Println("Enter -0 to stop.")
	for {
		fmt.Scanf("%d", &i)
		if i == -0 {
			break
		}
		is = append(is, i)
	}
	isf := removeOdd(is)
	fmt.Println(isf)
}

func removeOdd(is []int) []int {
	var isf []int
	for _, v := range is {
		if v%2 == 0 {
			isf = append(isf, v)
		}
	}

	return isf
}
