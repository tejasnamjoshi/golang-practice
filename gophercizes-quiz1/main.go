package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	qa := getQA()
	result := 0
	index := 1
	for q, a := range qa {
		fmt.Printf("Question %v: %v \n", index, q)
		s := ""
		fmt.Printf("Answer %v: ", index)
		fmt.Scanf("%s", &s)
		if s == a {
			result++
		}
		index++
		fmt.Println()
	}

	fmt.Println("Result is ", result)
}

func parseCsv() [][]string {
	fn := "problems.csv"
	if len(os.Args) >= 2 {
		fn = os.Args[1]
	}

	f, err := os.ReadFile(fn)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(f)))

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return records
}

func getQA() map[string]string {
	records := parseCsv()
	qa := map[string]string{}

	for _, r := range records {
		qa[r[0]] = r[1]
	}
	return qa
}
