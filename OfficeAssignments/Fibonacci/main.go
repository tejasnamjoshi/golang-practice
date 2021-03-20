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

	f := fibClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
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

// fibonacci is a function that returns
// a function that returns an int.
func fibClosure() func() int {
	a, b := 0, 1
	fmt.Printf("%d %d ", a, b)
	return func() int {
		a, b = b, a+b
		return b
	}
}
