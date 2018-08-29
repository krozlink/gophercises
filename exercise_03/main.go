package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	story, err := loadStory("gopher.json")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", serve(story))
	http.ListenAndServe(":80", nil)
}

func serve(s Story) http.HandlerFunc {
	t, err := template.ParseFiles("story.gohtml")
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if val, ok := s[path]; ok {
			t.ExecuteTemplate(w, "story", val)
		} else {
			http.NotFound(w, r)
		}
	}
}

func loadStory(filename string) (Story, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var story Story
	json.Unmarshal(contents, &story)
	return story, nil
}
