package main

import (
	"flag"
	"fmt"
	adventure "golang_practice/Gophercizes/Adventure"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the web app on")
	filename := flag.String("file", "gopher.json", "The JSON file for the Adventure")
	flag.Parse()
	fmt.Printf("Using the story %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	story, err := adventure.JSONStory(f)

	// tpl := template.Must(template.New("").Parse(storyTpl))
	// h := adventure.NewHandler(story, adventure.WithTemplate(tpl), adventure.WithPathFn(pathFn))
	h := adventure.NewHandler(story)

	fmt.Printf("starting server at %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/" || path == "/story" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTpl = `
<!DOCTYPE html>
<html lang="en">

<head>
    <title>Choose your own Adventure</title>
</head>

<body>
    <section class="page">
		<h1>{{.Title}}</h1>

		{{range .Paragraphs}}
		<p>{{.}}</p>
		{{end}}

		<ul>
			{{range .Options}}
			<li>
				<a href="/story/{{.Chapter}}">
					{{ .Text }}
				</a>
			</li>
			{{end}}
		</ul>
	</section>
	<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
</body>
</html>
`
