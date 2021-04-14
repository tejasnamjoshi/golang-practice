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

	/* f implements Writer interface */
	// f, err := os.Create("http_file2.html")
	// if err != nil {
	// 	panic(err)
	// }
	// io.Copy(f, resp.Body)

	lw := logWritter{}
	io.Copy(lw, resp.Body)
}

func (logWritter) Write(bs []byte) (int, error) {
	if err := os.WriteFile("http_file.html", bs, 0666); err != nil {
		fmt.Println("ERROR: ", err)
		return 0, err
	}

	return len(bs), nil
}
