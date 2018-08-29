package main

import (
	"flag"
	"fmt"
	cyoa "github.com/krozlink/gophercises/exercise_03"
	"net/http"
)

func main() {

	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	port := flag.Int("port", 3000, "The port to start the application on")

	story, err := cyoa.NewStory(*filename)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	http.ListenAndServe(fmt.Sprintf(":%v", *port), h)
}
