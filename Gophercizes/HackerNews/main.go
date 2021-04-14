package main

import (
	"flag"
	"fmt"
	"golang_practice/Gophercizes/HackerNews/hn"
	"html/template"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "Port", 3000, "Port number to start the application on.")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))
	http.HandleFunc("/", handler(tpl))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	// for i, id := range ids {
	// 	if i == 10 {
	// 		break
	// 	}
	// 	item, err := c.GetItemById(id)
	// 	if err != nil {
	// 		fmt.Println("Could not get item with Id: ", id)
	// 		continue
	// 	}
	// 	item.Print()
	// 	fmt.Println()
	// }
}

func handler(tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c hn.Client
		ids, err := c.GetTopStories()
		if err != nil {
			panic(err)
		}

	})
}
