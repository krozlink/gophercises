package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// Story represents a Choose Your Own Adventure story
// It consists of a collection of chapters that contain links
// to other chapters. The story starts at the chapter called "intro"
type Story map[string]Chapter

// Chapter represents a chapter within a Choose Your Own Adventure story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text    string `json:"text"`
		Chapter string `json:"arc"`
	} `json:"options"`
}

func NewStory(file string) (Story, error) {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var story Story
	json.Unmarshal(contents, &story)
	return story, nil
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tmpl = t
	}
}

func WithURLParser(p func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.parser = p
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{
		story:  s,
		tmpl:   template.Must(template.New("").Parse(defaultTemplate)),
		parser: defaultURLParser,
	}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func defaultURLParser(r *http.Request) string {
	path := r.URL.Path[1:]
	if path == "" || path == "/" {
		path = "intro"
	}

	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.parser(r)
	if val, ok := h.story[path]; ok {
		if err := h.tmpl.Execute(w, val); err != nil {
			fmt.Fprintln(w, err)
		}
	} else {
		http.NotFound(w, r)
	}
}

type handler struct {
	story  Story
	tmpl   *template.Template
	parser func(r *http.Request) string
}

var defaultTemplate = `
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
        <style>
            body {
                font-family: "Arial";
            }
        </style>
    </head>
    <body>

    <h1>{{.Title}}</h1>
    {{range .Story}}
        <p class="story">{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
        <li>
            <a href="/{{.Chapter}}">{{.Text}}</a>
        </li>
    {{end}}
    </ul>
    </body>
</head>
`
