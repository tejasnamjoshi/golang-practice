package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type quiz struct {
	question, answer string
}

func main() {
	qas := getQA()
	result := 0
	setTimer(&result, qas)

	for i, qa := range qas {
		showQuestion(qa.question, i)
		s := getAnswer(i)
		verifyAnswer(qa.answer, s, &result)

		fmt.Println()
	}

	showScore(result, len(qas))
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

func showScore(score int, total int) {
	fmt.Printf("You scored %d out of %d", score, total)
}

func showQuestion(question string, qNo int) {
	fmt.Printf("Question %v: %v \n", qNo+1, question)
}

func getAnswer(i int) string {
	s := ""
	fmt.Printf("Answer %v: ", i+1)
	fmt.Scan(&s)

	return s
}

func verifyAnswer(answer string, userAnswer string, result *int) {
	if userAnswer == answer {
		(*result)++
	}
}

func setTimer(result *int, qas []quiz) {
	fmt.Println("Press [Enter] to start test.")
	bufio.NewScanner(os.Stdout).Scan()
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("\nTime up!")
		showScore(*result, len(qas))
		os.Exit(0)
	})
}
