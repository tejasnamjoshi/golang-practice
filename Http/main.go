package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//  ! This is important. Read() is not written to write data to the slice if its already full. Hence create a blank slice with 9999 places.
	// bs := make([]byte, 9999)

	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWritter{}
	io.Copy(lw, resp.Body)
}

func (logWritter) Write(bs []byte) (int, error) {
	err := os.WriteFile("http_file.html", bs, 0666)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return 0, err
	}

	return len(bs), nil
}
