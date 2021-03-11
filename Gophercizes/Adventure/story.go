package adventure

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// ! struct names with first letter uppercase is automatically exported

func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

var tpl *template.Template

var defaultTemplate = `
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
				<a href="/{{.Chapter}}">
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

// HandlerOption is
type HandlerOption func(h *handler)

// WithTemplate is
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFn is
func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

// NewHandler is
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

// JSONStory is
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story

	if err := d.Decode(&story); err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	return story, nil
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found...", http.StatusNotFound)
}

// Story is
type Story map[string]Chapter

// Chapter is
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is
type Option struct {
	Chapter string `json:"arc"`
	Text    string `json:"text"`
}
