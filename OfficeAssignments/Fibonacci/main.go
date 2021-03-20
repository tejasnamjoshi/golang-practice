package main

import (
	"fmt"
	"os"
)

func main() {
	max := 100
	var input int
	fmt.Print("Enter limit : ")
	fmt.Scanf("%d", &input)

	if input > max {
		fmt.Printf("%d is greater than max value - %d", input, max)
		os.Exit(1)
	}

	fmt.Printf("\nWithout Recursion\n")
	fib(input)

	fmt.Printf("\n\nRecursion\n")
	fib(input)
}

func fib(input int) {
	a, b := 0, 1

	fmt.Printf("%d %d ", a, b)
	for i := 2; i < input; i++ {
		a, b = b, a+b
		fmt.Printf("%d ", b)
	}
}

func recFib(input int) int {
	if input < 1 {
		return 0
	}
	fmt.Println(input)

	return recFib(input-1) + recFib(input-2)
}
