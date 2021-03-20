package main

import "fmt"

func main() {
	var input int
	fmt.Println("Enter a number.")
	fmt.Scanf("%d", &input)

	if input%2 == 0 {
		fmt.Println("Number is even.")
	} else {
		fmt.Println("Number is odd.")
	}
}
