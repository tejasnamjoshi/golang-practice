package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type quiz struct {
	question, answer string
}

func main() {
	qas := getQA()
	result := 0
	for i, qa := range qas {
		fmt.Printf("Question %v: %v \n", i+1, qa.question)
		s := ""
		fmt.Printf("Answer %v: ", i+1)
		fmt.Scan(&s)
		if s == qa.answer {
			result++
		}
		fmt.Println()
	}

	fmt.Printf("You scored %d out of %d", result, len(qas))
}

func parseCsv() [][]string {
	fn := "problems.csv"
	if len(os.Args) >= 2 {
		fn = os.Args[1]
	}

	// f, err := os.ReadFile(fn)
	f, err := os.Open(fn)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	// r := csv.NewReader(strings.NewReader(string(f)))
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return records
}

func getQA() []quiz {
	records := parseCsv()
	qa := []quiz{}

	for _, r := range records {
		qa = append(qa, quiz{question: r[0], answer: r[1]})
	}

	return qa
}
