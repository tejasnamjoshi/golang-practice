package main

import (
	"fmt"
	"os"
)

func main() {
	var marks int
	fmt.Print("Enter marks : ")
	fmt.Scanf("%d", &marks)

	if marks > 100 {
		fmt.Println("Please enter marks less than 100 and try again.")
		os.Exit(1)
	}

	switch {
	case marks > 60:
		fmt.Println("Distintion")

	case marks > 40:
		fmt.Println("Pass")

	default:
		fmt.Println("Fail")
	}
}
