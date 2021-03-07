package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fn := getFileName()
	file, err := os.Open(fn)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
func getFileName() string {
	fn := os.Args
	if len(fn) < 2 {
		fmt.Println("No filename provided.")
		os.Exit(1)
	}

	return fn[1]
}
